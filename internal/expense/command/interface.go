package command

import (
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

// ExpenseCommandInterface ...
type ExpenseCommandInterface interface {
	// expense type
	CreateExpenseType(*model.ExpenseTypeCreate) error
	UpdateExpenseType(uuid.UUID, *model.ExpenseTypeUpdate) error
	DeleteExpenseType(uuid.UUID) error

	// expense
	CreateExpense(*sqlx.Tx, *model.ExpenseCreate) error

	// expense type image
	CreateExpenseImage(*sqlx.Tx, *model.ExpenseImageCreate) error
}
