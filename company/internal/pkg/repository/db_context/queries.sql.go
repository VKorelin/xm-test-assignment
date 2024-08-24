// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: queries.sql

package dbcontext

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM companies
WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteCompany, id)
	return err
}

const getCompany = `-- name: GetCompany :one
SELECT id, name, description, amount_of_employees, registered, type FROM companies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCompany(ctx context.Context, id pgtype.UUID) (Company, error) {
	row := q.db.QueryRow(ctx, getCompany, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.AmountOfEmployees,
		&i.Registered,
		&i.Type,
	)
	return i, err
}

const newCompany = `-- name: NewCompany :one
INSERT INTO companies (name, description, amount_of_employees, registered, type) VALUES ($1, $2, $3, $4, $5) RETURNING id
`

type NewCompanyParams struct {
	Name              string
	Description       pgtype.Text
	AmountOfEmployees int32
	Registered        bool
	Type              int32
}

func (q *Queries) NewCompany(ctx context.Context, arg NewCompanyParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, newCompany,
		arg.Name,
		arg.Description,
		arg.AmountOfEmployees,
		arg.Registered,
		arg.Type,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const updateCompany = `-- name: UpdateCompany :exec
UPDATE companies SET name = $1, description = $2, amount_of_employees = $3, registered = $4, type = $5 WHERE id = $6
`

type UpdateCompanyParams struct {
	Name              string
	Description       pgtype.Text
	AmountOfEmployees int32
	Registered        bool
	Type              int32
	ID                pgtype.UUID
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) error {
	_, err := q.db.Exec(ctx, updateCompany,
		arg.Name,
		arg.Description,
		arg.AmountOfEmployees,
		arg.Registered,
		arg.Type,
		arg.ID,
	)
	return err
}
