package domain

import "github.com/google/uuid"

type Customer struct {
	CustomerID uuid.UUID
	Name       string
}
