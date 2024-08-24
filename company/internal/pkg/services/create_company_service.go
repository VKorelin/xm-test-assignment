package services

import (
	"context"
	"xm/company/internal/pkg/models"
)

type CompanyCreator interface {
	Create(ctx context.Context, company *models.Company) (*models.Company, error)
}

type CreateCompanyService struct {
	companyCreator CompanyCreator
}

func NewCreateCompanyService(companyCreator CompanyCreator) *CreateCompanyService {
	return &CreateCompanyService{
		companyCreator: companyCreator,
	}
}

func (s *CreateCompanyService) Create(ctx context.Context, newCompany *models.Company) (*models.Company, error) {
	return s.companyCreator.Create(ctx, newCompany)
}
