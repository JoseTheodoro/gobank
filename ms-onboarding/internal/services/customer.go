package services

import pbCustomer "gobank/contracts/pb/customer"

type CustomerService struct {
	customerClient pbCustomer.CustomerClient
}

func NewCustomerService(c pbCustomer.CustomerClient) *CustomerService {
	return &CustomerService{
		customerClient: c,
	}
}
