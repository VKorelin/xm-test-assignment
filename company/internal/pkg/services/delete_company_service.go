package services

import (
	"context"
	"xm/company/internal/pkg/services/notifications"

	"github.com/gofrs/uuid"
)

type CompanyRemover interface {
	Delete(ctx context.Context, companyId uuid.UUID) error
}

type DeleteCompanyService struct {
	companyRemover      CompanyRemover
	notificationService notifications.NotificationService
}

func NewDeleteCompanyService(companyRemover CompanyRemover, notificationService notifications.NotificationService) *DeleteCompanyService {
	return &DeleteCompanyService{
		companyRemover:      companyRemover,
		notificationService: notificationService,
	}
}

func (s *DeleteCompanyService) Delete(ctx context.Context, companyId uuid.UUID) error {
	if err := s.companyRemover.Delete(ctx, companyId); err != nil {
		return err
	}

	return s.notificationService.Notify(ctx, companyId, notifications.Delete)
}
