package services

import (
	"context"
	"xm/company/internal/pkg/models"
)

type CompanyUpdater interface {
	Update(ctx context.Context, company *models.Company) error
}

type UpdateCompanyService struct {
	companyUpdater CompanyUpdater
}

func NewUpdateCompanyService(companyUpdater CompanyUpdater) *UpdateCompanyService {
	return &UpdateCompanyService{
		companyUpdater: companyUpdater,
	}
}

func (s *UpdateCompanyService) Update(ctx context.Context, company *models.Company) error {
	return s.companyUpdater.Update(ctx, company)
}
