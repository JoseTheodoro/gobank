package postgres

import (
	"context"
	"fmt"
	"gobank/ms-onboarding/internal/domain"
	"gobank/ms-onboarding/internal/infraestructure/database/postgres/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OnboardingPostgres struct {
	db    *pgxpool.Pool
	query *queries.Queries
}

func NewOnboardingPostgress(db *pgxpool.Pool) *OnboardingPostgres {
	return &OnboardingPostgres{
		db:    db,
		query: queries.New(db),
	}
}

func (o *OnboardingPostgres) Create(ctx context.Context, op domain.OnboardingProcess) (*domain.OnboardingProcess, error) {

	arg := queries.CreateOnboardingProcessParams{
		OnboardingID: op.OnboardingID,
		Email:        op.Email,
		Document:     op.Document,
		Status:       string(op.Status),
	}

	row, err := o.query.CreateOnboardingProcess(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("error on insert onboarding process > %w", err)
	}

	d, err := o.toDomain(row)
	if err != nil {
		return nil, fmt.Errorf("error on convert to onboarding domain > %w", err)
	}

	return d, nil
}

func (o *OnboardingPostgres) Update(ctx context.Context) error {
	return nil
}

func (o *OnboardingPostgres) toDomain(row queries.OnboardingProcess) (*domain.OnboardingProcess, error) {
	p := &domain.OnboardingProcess{
		ID:           row.ID,
		OnboardingID: row.OnboardingID,
		CustomerID:   row.CustomerID,
		AccountID:    row.AccountID,
		Email:        row.Email,
		Document:     row.Document,
		Status:       domain.OnboardingStatus(row.Status),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}

	return p, nil
}
