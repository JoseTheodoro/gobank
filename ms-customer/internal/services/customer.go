package services

import (
	"context"
	"gobank/ms-customer/internal/domain"

	"github.com/google/uuid"
)

type CustomerService struct {
	repository domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) *CustomerService {
	return &CustomerService{repository: repo}
}

func (s *CustomerService) CreateUser(ctx context.Context, input *domain.CustomerInput) (*domain.Customer, error) {
	c := &domain.Customer{
		CustomerID: uuid.New(),
		Name:       input.Name,
		Email:      input.Email,
		Document:   input.Document,
		Type:       input.Type,
		Status:     domain.UNDER_ANALYSIS,
	}
	return s.repository.CreateCustomer(ctx, c)
}
