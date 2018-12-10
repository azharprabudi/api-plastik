package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// ExpenseID ...
type ExpenseID struct {
	ExpenseID uuid.UUID `db:"id"`
}

// ExpenseRead ...
type ExpenseRead struct {
	ExpenseCreate
	ExpenseImageRead
}

// ExpenseCreate ...
type ExpenseCreate struct {
	ExpenseID     uuid.UUID `db:"id"`
	ExpenseTypeID int       `db:"expense_type_id"`
	Name          string    `db:"name"`
	Amount        float64   `db:"amount"`
	Note          string    `db:"note"`
	CreatedAt     time.Time `db:"created_at"`
}

// ExpenseUpdate ...
type ExpenseUpdate struct {
	Name          string  `db:"name"`
	ExpenseTypeID int     `db:"expense_type_id"`
	Amount        float64 `db:"amount"`
	Note          string  `db:"note"`
}
