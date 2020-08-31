package schemas

import (
	"github.com/satori/go.uuid"
)

// FinanceTransactionCreate ...
type FinanceTransactionCreate struct {
	Name             string    `validate:"required"`
	Description      string    `validate:"omitempty"`
	Amount           float64   `validate:"required,gt=0"`
	FinanceAccountID uuid.UUID `validate:"required" form:"financeAccountId" `
}

// FinanceTransactionUpdate ...
type FinanceTransactionUpdate struct {
	Name        string  `validate:"required"`
	Description string  `validate:"omitempty"`
	Amount      float64 `validate:"required,gt=0"`
}

// FinanceTransactionURI is param uri validation for id
type FinanceTransactionURI struct {
	ID               uuid.UUID `uri:"id" binding:"required"`
	FinanceAccountID uuid.UUID `uri:"financeAccountId" binding:"required"`
}

// Summary is param uri validation for id
type Summary struct {
	Month int `validate:"omitempty,numeric" form:"month"`
	Day   int `validate:"omitempty,numeric" form:"day"`
}
