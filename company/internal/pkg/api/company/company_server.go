package company

import (
	"context"
	"xm/company/internal/pkg/api/company/converters"
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

type UpdateService interface {
	Update(ctx context.Context, company *models.Company) error
}

type DeleteService interface {
	Delete(ctx context.Context, companyId uuid.UUID) error
}

var _ servicepb.CompanyServiceServer = (*CompanyServiceServerImpl)(nil)

type CompanyServiceServerImpl struct {
	servicepb.UnimplementedCompanyServiceServer
	fetchService  FetchService
	createService CreateService
	updateService UpdateService
	deleteService DeleteService
}

func NewCompanyServerImpl(fetchService FetchService, createService CreateService, updateService UpdateService, deleteService DeleteService) *CompanyServiceServerImpl {
	return &CompanyServiceServerImpl{
		fetchService:  fetchService,
		createService: createService,
		updateService: updateService,
		deleteService: deleteService,
	}
}

func (s *CompanyServiceServerImpl) Get(ctx context.Context, request *servicepb.GetCompanyRequest) (*servicepb.GetCompanyResponse, error) {

	companyId, _ := uuid.FromString(request.CompanyId)

	company, err := s.fetchService.Fetch(ctx, companyId)
	if err != nil {
		return nil, NotFound(err)
	}

	return &servicepb.GetCompanyResponse{
		Company: converters.ConvertModelToProto(company),
	}, nil
}

func (s *CompanyServiceServerImpl) Patch(ctx context.Context, request *servicepb.PatchCompanyRequest) (*emptypb.Empty, error) {
	err := s.updateService.Update(ctx, converters.ConvertProtoToModel(request.Company))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func (s *CompanyServiceServerImpl) Delete(ctx context.Context, request *servicepb.DeleteCompanyRequest) (*emptypb.Empty, error) {
	companyId, _ := uuid.FromString(request.CompanyId)

	err := s.deleteService.Delete(ctx, companyId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, err
}

func (s *CompanyServiceServerImpl) Create(ctx context.Context, request *servicepb.CreateCompanyRequest) (*servicepb.CreateCompanyResponse, error) {
	company, err := s.createService.Create(ctx, converters.ConvertCreateCompanyRequestToModel(request))

	if err != nil {
		return nil, InternalError(err)
	}

	return &servicepb.CreateCompanyResponse{
		Company: converters.ConvertModelToProto(company),
	}, nil
}
