package services

import (
	"context"
	"xm/company/internal/pkg/models"

	"github.com/google/uuid"
)

type CompanyProvider interface {
	Get(ctx context.Context, companyId uuid.UUID) (*models.Company, error)
}

type FetchCompanyService struct {
	companyProvider CompanyProvider
}

func NewFetchCompanyService(companyProvider CompanyProvider) *FetchCompanyService {
	return &FetchCompanyService{
		companyProvider: companyProvider,
	}
}

func (s *FetchCompanyService) Fetch(ctx context.Context, companyId uuid.UUID) (*models.Company, error) {
	return s.companyProvider.Get(ctx, companyId)
}
