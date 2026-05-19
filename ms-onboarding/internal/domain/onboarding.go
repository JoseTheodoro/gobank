package domain

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

type OnboardingProcess struct {
	ID           int64            `json:"id"`
	OnboardingID uuid.UUID        `json:"onboarding_id"`
	CustomerID   *uuid.UUID       `json:"customer_id"`
	AccountID    *uuid.UUID       `json:"account_id"`
	Email        string           `json:"email"`
	Document     string           `json:"document"`
	Status       OnboardingStatus `json:"status"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    *time.Time       `json:"updated_at"`
}

type StartOnboardingInput struct {
	Customer    CustomerInput
	Credentials CredentialsInput
	Device      DeviceInput
}

type CredentialsInput struct {
	Email    string
	Password string
}
type CustomerInput struct {
	Name     string
	Document string
	Type     string
}
type DeviceInput struct {
	IPAddr    string
	UserAgent string
	DeviceID  string
}

type OnboardingProcessHistory struct{}
