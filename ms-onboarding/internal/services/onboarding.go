package services

import (
	"context"
	"fmt"
	"gobank/ms-onboarding/internal/domain"

	"github.com/google/uuid"
)

type OnboardingService struct {
	CustomerService *CustomerService
	AuthService     *AuthService
	AccountService  *AccountService
	KYCService      *KYCService
	Repository      domain.Repository
}

func NewOnboardingService(c *CustomerService, a *AuthService, acc *AccountService, k *KYCService, r domain.Repository) *OnboardingService {
	return &OnboardingService{
		CustomerService: c,
		AuthService:     a,
		AccountService:  acc,
		KYCService:      k,
		Repository:      r,
	}
}

func (o *OnboardingService) Run(ctx context.Context, input domain.StartOnboardingInput) error {

	//create onboarding status STARTED.
	onboardingStated := domain.OnboardingProcess{
		OnboardingID: uuid.New(),
		Email:        input.Credentials.Email,
		Document:     input.Customer.Document,
		Status:       domain.Started,
	}
	onboardingProcess, err := o.Repository.Create(ctx, onboardingStated)
	if err != nil {
		return fmt.Errorf("error on starting onboarding process > %w ", err)
	}
	fmt.Println("[RUN]", "onboardingprocess=", onboardingProcess)
	fmt.Println("ready to request ms-customer")

	// call ms-customer
	customer, err := o.CustomerService.CreateCustomer(ctx, input)
	if err != nil {
		return fmt.Errorf("error on creating customer > %w ", err)
	}
	fmt.Println("customer=", customer)
	// update onboarding status to CUSTOMER_CREATED
	if err := o.Repository.SetCustomer(ctx, onboardingProcess, customer, domain.CustomerCreated); err != nil {
		return fmt.Errorf("error on set customer > %w ", err)
	}
	// call ms-auth
	o.AuthService.CreateCredentials(ctx, input.Credentials, customer)
	// update onboarding status to CREDENTIAL_CREATED
	// call ms-account
	// update onboarding status to ACCOUNT_CREATED
	// call ms-kyc
	// updated onboarding status to AWAITING_KYC

	return nil
}
