package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Expense ...
type Expense struct {
	Name string `json:"name"`
}

// ExpenseTypeReq ...
type ExpenseTypeReq struct {
	Expense
}

// ExpenseTypeRes ...
type ExpenseTypeRes struct {
	Expense
	ExpenseTypeID uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
}
