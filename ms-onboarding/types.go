package main

import (
	"time"

	"github.com/google/uuid"
)

type OnboardingStatus string

const (
	Started            OnboardingStatus = "STARTED"
	CustomerCreated    OnboardingStatus = "CUSTOMER_CREATED"
	CredentialsCreated OnboardingStatus = "CREDENTIALS_CREATED"
	AccountCreated     OnboardingStatus = "ACCOUNT_CREATED"
	KYCAWaiting        OnboardingStatus = "AWAITING_KYC"
	CompletedApproved  OnboardingStatus = "COMPLETED_APPROVED"
	CompledtedRejected OnboardingStatus = "COMPLETED_REJECTED"
	FailedCompensated  OnboardingStatus = "FAILED_COMPENSATED"
)

type Onboarding struct {
	ID           int64            `json:"id"`
	OnboardingID uuid.UUID        `json:"onboarding_id"`
	CustomerID   *uuid.UUID       `json:"customer_id"`
	AccountID    *uuid.UUID       `json:"account_id"`
	Status       OnboardingStatus `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    *time.Time       `json:"updated_at"`
}
