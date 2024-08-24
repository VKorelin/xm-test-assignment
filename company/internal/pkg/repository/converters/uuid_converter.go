package converters

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertToPgUuid(uuid uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: uuid,
		Valid: true,
	}
}

func ConvertFromPgUuid(pgUuid pgtype.UUID) (uuid.UUID, error) {
	return uuid.FromBytes(pgUuid.Bytes[:])
}
