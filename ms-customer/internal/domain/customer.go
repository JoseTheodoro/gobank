package domain

import (
	"time"

	"github.com/google/uuid"
)

type CustomerStatus string

const (
	UNDER_ANALYSIS CustomerStatus = "UNDER_ANALYSIS"
	APPROVED       CustomerStatus = "APPROVED"
	BANNED         CustomerStatus = "BANNED"
)

type CustomerType string

const (
	INDIVIDUAL CustomerType = "INDIVIDUAL"
	BUSINESS   CustomerType = "BUSINESS"
	SYSTEM     CustomerType = "SYSTEM"
)

type Customer struct {
	ID         int64
	CustomerID uuid.UUID
	Name       string
	Email      string
	Document   string
	Type       CustomerType
	Status     CustomerStatus
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type CustomerInput struct {
	Name     string
	Email    string
	Document string
	Type     CustomerType
}
