package services

import pbAuth "gobank/contracts/pb/auth"

type AuthService struct {
	authClient pbAuth.AuthClient
}

func NewAuthService(c pbAuth.AuthClient) *AuthService {
	return &AuthService{
		authClient: c,
	}
}
