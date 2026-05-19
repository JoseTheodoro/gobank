package services

import pbAccount "gobank/contracts/pb/account"

type AccountService struct {
	accountClient pbAccount.AccountClient
}

func NewAccountService(c pbAccount.AccountClient) *AccountService {
	return &AccountService{accountClient: c}
}
