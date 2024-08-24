package services

import (
	"context"

	"github.com/gofrs/uuid"
)

type CompanyRemover interface {
	Delete(ctx context.Context, companyId uuid.UUID) error
}

type DeleteCompanyService struct {
	companyRemover CompanyRemover
}

func NewDeleteCompanyService(companyRemover CompanyRemover) *DeleteCompanyService {
	return &DeleteCompanyService{
		companyRemover: companyRemover,
	}
}

func (s *DeleteCompanyService) Delete(ctx context.Context, companyId uuid.UUID) error {
	return s.companyRemover.Delete(ctx, companyId)
}
