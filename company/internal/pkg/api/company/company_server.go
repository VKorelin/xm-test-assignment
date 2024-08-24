package company

import (
	"context"
	"xm/company/internal/pkg/models"
	servicepb "xm/company/pkg/api/company/v1"

	"github.com/google/uuid"
)

type FetchServie interface {
	Fetch(ctx context.Context, companyId uuid.UUID) (*models.Company, error)
}

var _ servicepb.CompanyServiceServer = (*CompanyServiceServerImpl)(nil)

type CompanyServiceServerImpl struct {
	servicepb.UnimplementedCompanyServiceServer
	fetchService FetchServie
}

func NewOrderServerImpl(fetchService FetchServie) *CompanyServiceServerImpl {
	return &CompanyServiceServerImpl{
		fetchService: fetchService,
	}
}

func (s *CompanyServiceServerImpl) Get(ctx context.Context, request *servicepb.GetCompanyRequest) (*servicepb.GetCompanyResponse, error) {

	companyId, _ := uuid.Parse(request.CompanyId)

	company, err := s.fetchService.Fetch(ctx, companyId)
	if err != nil {
		return nil, NotFound(err)
	}

	return &servicepb.GetCompanyResponse{
		Company: &servicepb.Company{
			Id:                company.Id.String(),
			Name:              company.Name,
			Description:       company.Description,
			AmountOfEmployees: int32(company.AmountOfEmployees),
			Registered:        company.Registered,
			Type:              servicepb.CompanyType(company.Type),
		},
	}, nil
}
