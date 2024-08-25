package services

import (
	"context"
	"xm/company/internal/pkg/models"
	"xm/company/internal/pkg/services/notifications"
)

type CompanyCreator interface {
	Create(ctx context.Context, company *models.Company) (*models.Company, error)
}

type CreateCompanyService struct {
	companyCreator      CompanyCreator
	notificationService notifications.NotificationService
}

func NewCreateCompanyService(companyCreator CompanyCreator, notificationService notifications.NotificationService) *CreateCompanyService {
	return &CreateCompanyService{
		companyCreator:      companyCreator,
		notificationService: notificationService,
	}
}

func (s *CreateCompanyService) Create(ctx context.Context, newCompany *models.Company) (*models.Company, error) {
	company, err := s.companyCreator.Create(ctx, newCompany)
	if err != nil {
		return nil, err
	}

	if err := s.notificationService.Notify(ctx, company.Id, notifications.Create); err != nil {
		return nil, err
	}

	return company, err
}
