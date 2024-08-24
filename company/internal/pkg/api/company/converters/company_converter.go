package converters

import (
	"xm/company/internal/pkg/models"
	servicepb "xm/company/pkg/api/company/v1"

	"github.com/gofrs/uuid"
)

func ConvertModelToProto(company *models.Company) *servicepb.Company {
	return &servicepb.Company{
		Id:                company.Id.String(),
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              servicepb.CompanyType(company.Type),
	}
}

func ConvertProtoToModel(company *servicepb.Company) *models.Company {
	companyId, _ := uuid.FromString(company.Id)

	return &models.Company{
		Id:                companyId,
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              models.CompanyType(company.Type),
	}
}

func ConvertCreateCompanyRequestToModel(company *servicepb.CreateCompanyRequest) *models.Company {
	return &models.Company{
		Name:              company.Name,
		Description:       company.Description,
		AmountOfEmployees: company.AmountOfEmployees,
		Registered:        company.Registered,
		Type:              models.CompanyType(company.Type),
	}
}
