package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// Expense ...
type Expense struct {
	ExpenseID     uuid.UUID `db:"id"`
	ExpenseTypeID uuid.UUID `db:"expense_type_id"`
	Name          string    `db:"name"`
	Amount        float64   `db:"amount"`
	Note          string    `db:"note"`
	CreatedAt     time.Time `db:"created_at"`
}

// ExpenseRead ...
type ExpenseRead struct {
	Expense
}

// ExpenseReadDetail ...
type ExpenseReadDetail struct {
	Expense
	ExpenseImage
}

// ExpenseCreate ...
type ExpenseCreate struct {
	Expense
}
