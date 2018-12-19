package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ExpenseImage ...
type ExpenseImage struct {
	ExpenseImageID uuid.UUID `db:"id"`
	ExpenseID      uuid.UUID `db:"expense_id"`
	Image          string    `db:"image"`
	CreatedAt      time.Time `db:"created_at"`
}

// ExpenseImageRead ...
type ExpenseImageRead struct {
	ExpenseImage
}

// ExpenseImageCreate ...
type ExpenseImageCreate struct {
	ExpenseImage
}
