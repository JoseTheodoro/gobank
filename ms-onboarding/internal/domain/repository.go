package domain

import "context"

type Repository interface {
	Create(ctx context.Context, process OnboardingProcess) (*OnboardingProcess, error)
	Update(ctx context.Context) error
}
