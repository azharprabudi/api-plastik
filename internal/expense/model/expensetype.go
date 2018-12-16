package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// ExpenseType ...
type ExpenseType struct {
	ExpenseTypeID uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	CreatedAt     time.Time `db:"created_at"`
}

// ExpenseTypeRead ...
type ExpenseTypeRead struct {
	ExpenseType
}

// ExpenseTypeCreate ...
type ExpenseTypeCreate struct {
	ExpenseType
}

// ExpenseTypeUpdate ...
type ExpenseTypeUpdate struct {
	Name string `db:"name"`
}
