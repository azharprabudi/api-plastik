package command

import (
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
)

// ExpenseCommandInterface ...
type ExpenseCommandInterface interface {
	// expense type
	CreateExpenseType(*model.ExpenseTypeCreate) (int64, error)
	UpdateExpenseType(int, *model.ExpenseTypeUpdate) error
	DeleteExpenseType(int) error

	// expense
	CreateExpense(*sqlx.Tx, *model.ExpenseCreate) error

	// expense type image
	CreateExpenseImage(*sqlx.Tx, *model.ExpenseImageCreate) error
}
