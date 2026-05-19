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

func (o *OnboardingPostgres) SetCustomer(ctx context.Context, onboarding *domain.OnboardingProcess, customer *domain.Customer, newStatus domain.OnboardingStatus) error {

	c := queries.SetCustomerOnboardingProcessParams{
		CustomerID: &customer.CustomerID,
		Status:     string(newStatus),
		ID:         onboarding.ID,
	}

	err := o.query.SetCustomerOnboardingProcess(ctx, c)
	if err != nil {
		return fmt.Errorf("error on setting customer on onboarding process > %w", err)
	}
	return nil
}

func (o *OnboardingPostgres) SetAccount(ctx context.Context, onboarding *domain.OnboardingProcess, account *domain.Account, newStatus domain.OnboardingStatus) error {

	a := queries.SetAccountOnboardingProcessParams{
		AccountID: account.AccountID,
		Status:    string(newStatus),
		ID:        onboarding.ID,
	}

	err := o.query.SetAccountOnboardingProcess(ctx, a)
	if err != nil {
		return fmt.Errorf("error on set account on onboarding process > %w", err)
	}

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
