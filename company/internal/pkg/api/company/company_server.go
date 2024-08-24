package company

import (
	"context"
	"xm/company/internal/pkg/models"
	servicepb "xm/company/pkg/api/company/v1"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FetchService interface {
	Fetch(ctx context.Context, companyId uuid.UUID) (*models.Company, error)
}

type CreateService interface {
	Create(ctx context.Context, newCompany *models.Company) (*models.Company, error)
}

var _ servicepb.CompanyServiceServer = (*CompanyServiceServerImpl)(nil)

type CompanyServiceServerImpl struct {
	servicepb.UnimplementedCompanyServiceServer
	fetchService  FetchService
	createService CreateService
}

func NewCompanyServerImpl(fetchService FetchService, createService CreateService) *CompanyServiceServerImpl {
	return &CompanyServiceServerImpl{
		fetchService:  fetchService,
		createService: createService,
	}
}

func (s *CompanyServiceServerImpl) Get(ctx context.Context, request *servicepb.GetCompanyRequest) (*servicepb.GetCompanyResponse, error) {

	companyId, _ := uuid.FromString(request.CompanyId)

	company, err := s.fetchService.Fetch(ctx, companyId)
	if err != nil {
		return nil, NotFound(err)
	}

	return &servicepb.GetCompanyResponse{
		Company: &servicepb.Company{
			Id:                company.Id.String(),
			Name:              company.Name,
			Description:       company.Description,
			AmountOfEmployees: company.AmountOfEmployees,
			Registered:        company.Registered,
			Type:              servicepb.CompanyType(company.Type),
		},
	}, nil
}

func (s *CompanyServiceServerImpl) Patch(ctx context.Context, request *servicepb.PatchCompanyRequest) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

func (s *CompanyServiceServerImpl) Delete(ctx context.Context, request *servicepb.DeleteCompanyRequest) (*emptypb.Empty, error) {

	return &emptypb.Empty{}, nil
}

func (s *CompanyServiceServerImpl) Create(ctx context.Context, request *servicepb.CreateCompanyRequest) (*servicepb.CreateCompanyResponse, error) {
	company, err := s.createService.Create(ctx, &models.Company{
		Name:              request.Name,
		Description:       request.Description,
		AmountOfEmployees: request.AmountOfEmployees,
		Registered:        request.Registered,
		Type:              models.CompanyType(request.Type),
	})

	if err != nil {
		return nil, InternalError(err)
	}

	return &servicepb.CreateCompanyResponse{
		Company: &servicepb.Company{
			Id:                company.Id.String(),
			Name:              company.Name,
			Description:       company.Description,
			AmountOfEmployees: company.AmountOfEmployees,
			Registered:        company.Registered,
			Type:              servicepb.CompanyType(company.Type),
		},
	}, nil
}
