// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package dbcontext

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Company struct {
	ID                pgtype.UUID
	Name              string
	Description       pgtype.Text
	AmountOfEmployees int32
	Registered        bool
	Type              int32
}
