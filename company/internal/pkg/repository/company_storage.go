package repository

import (
	"context"
	dbcontext "xm/company/internal/pkg/repository/db_context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyStorage struct {
	dbPool *pgxpool.Pool
}

func NewCompanyStorage(dbPool *pgxpool.Pool) *CompanyStorage {
	return &CompanyStorage{
		dbPool: dbPool,
	}
}

func (f *CompanyStorage) CreateDbContext(ctx context.Context) dbcontext.Querier {
	return dbcontext.New(f.dbPool)
}
