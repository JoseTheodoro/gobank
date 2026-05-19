package handler

import (
	"context"
	"fmt"
	pb "gobank/contracts/pb/customer"
	"gobank/ms-customer/internal/domain"
	"gobank/ms-customer/internal/services"
)

type Handle struct {
	pb.UnimplementedCustomerServer
	customerService *services.CustomerService
}

func NewHandle(c *services.CustomerService) *Handle {
	return &Handle{
		customerService: c,
	}
}

func (s *Handle) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	fmt.Printf("request received: %s", in.String())

	customerInput := &domain.CustomerInput{
		Name:     in.GetName(),
		Email:    in.GetEmail(),
		Document: in.GetDocument(),
		Type:     domain.CustomerType(in.GetType()),
	}

	customer, err := s.customerService.CreateUser(ctx, customerInput)
	if err != nil {
		return nil, err
	}

	return &pb.CreateCustomerResponse{
		CustomerId: customer.CustomerID.String(),
	}, nil
}
