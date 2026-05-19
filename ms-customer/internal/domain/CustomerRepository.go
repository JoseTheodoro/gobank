package domain

import "context"

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, cutomer *Customer) (*Customer, error)
}
