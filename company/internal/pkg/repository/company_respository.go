package repository

import (
	"context"
	"xm/company/internal/pkg/models"

	"github.com/google/uuid"
)

type CompanyRepository struct {
}

func NewOrderRepository() *CompanyRepository {
	return &CompanyRepository{}
}

func (r *CompanyRepository) Get(ctx context.Context, orderId uuid.UUID) (*models.Company, error) {
	return &models.Company{
		Id:                uuid.New(),
		Name:              "test",
		Description:       "test description",
		AmountOfEmployees: 12,
		Registered:        true,
		Type:              models.Corporations,
	}, nil
}
