package notifications

import (
	"context"

	"github.com/gofrs/uuid"
)

type NotificationService interface {
	Notify(ctx context.Context, id uuid.UUID, changeType ChangeType) error
}
