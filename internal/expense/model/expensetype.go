package model

import (
	"time"
)

// ExpenseTypeID ...
type ExpenseTypeID struct {
	ExpenseTypeID int `db:"id"`
}

// ExpenseTypeModelRead ...
type ExpenseTypeModelRead struct {
	ExpenseTypeID
	ExpenseTypeCreate
}

// ExpenseTypeCreate ...
type ExpenseTypeCreate struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

// ExpenseTypeUpdate ...
type ExpenseTypeUpdate struct {
	Name string `db:"name"`
}
