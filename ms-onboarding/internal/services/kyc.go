package services

import pbKYC "gobank/contracts/pb/kyc"

type KYCService struct {
	KYCClient pbKYC.KYCClient
}

func NewKYCService(c pbKYC.KYCClient) *KYCService {
	return &KYCService{
		KYCClient: c,
	}
}
