package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ExpenseReq ...
type ExpenseReq struct {
	ExpenseTypeID uuid.UUID `json:"expenseTypeId"`
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	Note          string    `json:"note"`
	Images        []string  `json:"images"`
}

// ExpenseRes ...
type ExpenseRes struct {
	ExpenseID     uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	ExpenseTypeID uuid.UUID `json:"expenseTypeId"`
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	Note          string    `json:"note"`
}

// ExpenseResDetail ...
type ExpenseResDetail struct {
	ExpenseID     uuid.UUID         `json:"id"`
	CreatedAt     time.Time         `json:"createdAt"`
	ExpenseTypeID uuid.UUID         `json:"expenseTypeId"`
	Name          string            `json:"name"`
	Amount        float64           `json:"amount"`
	Note          string            `json:"note"`
	Images        []ExpenseImageRes `json:"images"`
}
