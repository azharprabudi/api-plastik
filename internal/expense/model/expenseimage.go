package model

import (
	"github.com/satori/go.uuid"
)

// ExpenseImage ...
type ExpenseImage struct {
	ExpenseImageID uuid.UUID `db:"id"`
	ExepenseID     uuid.UUID `db:"expense_id"`
	Image          string    `db:"image"`
}

// ExpenseImageRead ...
type ExpenseImageRead struct {
	ExpenseImage
}

// ExpenseImageCreate ...
type ExpenseImageCreate struct {
	ExpenseImage
}
