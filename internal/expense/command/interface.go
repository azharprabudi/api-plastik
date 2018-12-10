package command

import (
	"github.com/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// ExpenseCommandInterface ...
type ExpenseCommandInterface interface {
	// expense type
	CreateExpenseType(*model.ExpenseTypeCreate) (int64, error)
	UpdateExpenseType(int, *model.ExpenseTypeUpdate) error
	DeleteExpenseType(int) error

	// expense
	CreateExpense(*sqlx.Tx, *model.ExpenseCreate) error
	UpdateExpense(*sqlx.Tx, string, *model.ExpenseUpdate) error
	DeleteExpense(*sqlx.Tx, string) error

	// expense type image
	CreateExpenseImage(*sqlx.Tx, *model.ExpenseImageCreate) error
	DeleteExpenseImage(*sqlx.Tx, uuid.UUID) error
}
