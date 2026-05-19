package services

import (
	"context"
	pbAuth "gobank/contracts/pb/auth"
	"gobank/ms-onboarding/internal/domain"
)

type AuthService struct {
	authClient pbAuth.AuthClient
}

func NewAuthService(c pbAuth.AuthClient) *AuthService {
	return &AuthService{
		authClient: c,
	}
}

func (a *AuthService) CreateCredentials(ctx context.Context, input domain.CredentialsInput, c *domain.Customer) error {

	return nil
}
