package services

import (
	"context"
	"fmt"
	pbCustomer "gobank/contracts/pb/customer"
	"gobank/ms-onboarding/internal/domain"

	"github.com/google/uuid"
)

type CustomerService struct {
	customerClient pbCustomer.CustomerClient
}

func NewCustomerService(c pbCustomer.CustomerClient) *CustomerService {
	return &CustomerService{
		customerClient: c,
	}
}

func (c *CustomerService) CreateCustomer(ctx context.Context, input domain.CustomerInput) (*domain.Customer, error) {

	createCustomerRequest := pbCustomer.CreateCustomerRequest{Name: input.Name}

	response, err := c.customerClient.CreateCustomer(ctx, &createCustomerRequest)
	if err != nil {
		return nil, fmt.Errorf("error on service create customer > %w", err)
	}

	customer := &domain.Customer{
		CustomerID: uuid.New(),
		Name:       response.GetB(),
	}

	return customer, nil
}
