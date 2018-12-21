package command

import (
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// ExpenseCommandInterface ...
type ExpenseCommandInterface interface {
	CreateExpenseType(*model.ExpenseTypeCreate) error
	UpdateExpenseType(uuid.UUID, *model.ExpenseTypeUpdate) error
	DeleteExpenseType(uuid.UUID) error
	CreateExpense(*sqlx.Tx, *model.ExpenseCreate) error
	CreateExpenseImage(*sqlx.Tx, *model.ExpenseImageCreate) error
}
