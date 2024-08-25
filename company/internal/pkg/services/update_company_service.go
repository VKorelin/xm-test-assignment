package services

import (
	"context"
	"xm/company/internal/pkg/models"
	"xm/company/internal/pkg/services/notifications"
)

type CompanyUpdater interface {
	Update(ctx context.Context, company *models.Company) error
}

type UpdateCompanyService struct {
	companyUpdater      CompanyUpdater
	notificationService notifications.NotificationService
}

func NewUpdateCompanyService(companyUpdater CompanyUpdater, notificationService notifications.NotificationService) *UpdateCompanyService {
	return &UpdateCompanyService{
		companyUpdater:      companyUpdater,
		notificationService: notificationService,
	}
}

func (s *UpdateCompanyService) Update(ctx context.Context, company *models.Company) error {
	if err := s.companyUpdater.Update(ctx, company); err != nil {
		return err
	}

	return s.notificationService.Notify(ctx, company.Id, notifications.Patch)
}
