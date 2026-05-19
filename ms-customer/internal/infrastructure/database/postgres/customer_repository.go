package postgres

import (
	"context"
	"fmt"
	"gobank/ms-customer/internal/domain"
	"gobank/ms-customer/internal/infrastructure/database/postgres/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepositoryPostgres struct {
	db    *pgxpool.Pool
	query *queries.Queries
}

func NewCustomerRepositoryPostgres(db *pgxpool.Pool) *CustomerRepositoryPostgres {
	return &CustomerRepositoryPostgres{
		db:    db,
		query: queries.New(db),
	}
}

func (r *CustomerRepositoryPostgres) CreateCustomer(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {

	arg := queries.CreateParams{
		CustomerID: c.CustomerID,
		Name:       c.Name,
		Email:      c.Email,
		Document:   c.Document,
		Type:       string(c.Type),
		Status:     string(c.Status),
	}

	row, err := r.query.Create(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("error on insert customer on database > %w", err)
	}

	customer := toDomain(&row)

	return customer, nil

}

func toDomain(c *queries.Customer) *domain.Customer {
	return &domain.Customer{
		ID:         c.ID,
		CustomerID: c.CustomerID,
		Name:       c.Name,
		Email:      c.Email,
		Document:   c.Document,
		Type:       domain.CustomerType(c.Type),
		Status:     domain.CustomerStatus(c.Status),
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}
