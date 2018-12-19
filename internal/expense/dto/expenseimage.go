package dto

import (
	uuid "github.com/satori/go.uuid"
)

// ExpenseImageRes ...
type ExpenseImageRes struct {
	ExpenseImageID uuid.UUID `json:"id"`
	ExpenseID      uuid.UUID `json:"expenseId"`
	Image          string    `json:"image"`
}
