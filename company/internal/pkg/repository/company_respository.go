package repository

import (
	"context"
	"xm/company/internal/pkg/models"
	"xm/company/internal/pkg/repository/converters"
	dbcontext "xm/company/internal/pkg/repository/db_context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type CompanyRepository struct {
	storage Storage
	logger  *zap.Logger
}

type Storage interface {
	CreateDbContext(ctx context.Context) dbcontext.Querier
}

func NewCompanyRepository(storage Storage, logger *zap.Logger) *CompanyRepository {
	return &CompanyRepository{
		storage: storage,
		logger:  logger,
	}
}

func (r *CompanyRepository) Get(ctx context.Context, companyId uuid.UUID) (*models.Company, error) {

	context := r.storage.CreateDbContext(ctx)

	company, err := context.GetCompany(ctx, converters.ConvertToPgUuid(companyId))

	if err != nil {
		return nil, err
	}

	return &models.Company{
		Id:                companyId,
		Name:              company.Name,
		Description:       company.Description.String,
		AmountOfEmployees: uint32(company.AmountOfEmployees),
		Registered:        company.Registered,
		Type:              models.CompanyType(company.Type),
	}, nil
}

func (r *CompanyRepository) Create(ctx context.Context, newCompany *models.Company) (*models.Company, error) {

	context := r.storage.CreateDbContext(ctx)

	pgCompanyId, err := context.NewCompany(ctx, dbcontext.NewCompanyParams{
		Name: newCompany.Name,
		Description: pgtype.Text{
			String: newCompany.Description,
			Valid:  true},
		AmountOfEmployees: int32(newCompany.AmountOfEmployees),
		Registered:        newCompany.Registered,
		Type:              int32(newCompany.Type),
	})

	if err != nil {
		r.logger.Error("Could not create new company", zap.String("name", newCompany.Name), zap.Error(err))
		return nil, err
	}

	companyId, err := converters.ConvertFromPgUuid(pgCompanyId)
	if err != nil {
		r.logger.Error("Could not convert pg company ID to UUID", zap.Error(err))
		return nil, err
	}

	return &models.Company{
		Id:                companyId,
		Name:              newCompany.Name,
		Description:       newCompany.Description,
		AmountOfEmployees: newCompany.AmountOfEmployees,
		Registered:        newCompany.Registered,
		Type:              newCompany.Type,
	}, nil
}

func (r *CompanyRepository) Update(ctx context.Context, company *models.Company) error {

	context := r.storage.CreateDbContext(ctx)

	err := context.UpdateCompany(ctx, dbcontext.UpdateCompanyParams{
		Name: company.Name,
		Description: pgtype.Text{
			String: company.Description,
			Valid:  true},
		AmountOfEmployees: int32(company.AmountOfEmployees),
		Registered:        company.Registered,
		Type:              int32(company.Type),
		ID:                converters.ConvertToPgUuid(company.Id),
	})

	if err != nil {
		r.logger.Error("Could not update company", zap.String("ID", company.Id.String()), zap.Error(err))
	}

	return err
}

func (r *CompanyRepository) Delete(ctx context.Context, companyId uuid.UUID) error {

	context := r.storage.CreateDbContext(ctx)

	err := context.DeleteCompany(ctx, converters.ConvertToPgUuid(companyId))

	if err != nil {
		r.logger.Error("Could not delete company", zap.String("ID", companyId.String()), zap.Error(err))
	}

	return err
}
