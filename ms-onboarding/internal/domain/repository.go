package domain

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, process OnboardingProcess) (*OnboardingProcess, error)
	SetCustomer(ctx context.Context, onboarding *OnboardingProcess, cutomer *Customer, newStatus OnboardingStatus) error
	SetAccount(ctx context.Context, onboarding *OnboardingProcess, account *Account, newStatus OnboardingStatus) error
}
