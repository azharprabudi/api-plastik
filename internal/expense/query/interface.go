package query

import (
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	uuid "github.com/satori/go.uuid"
)

// ExpenseQueryInterface ...
type ExpenseQueryInterface interface {
	// Expense Type
	GetExpenseType() ([]*model.ExpenseTypeRead, error)
	GetExpenseTypeByID(uuid.UUID) *model.ExpenseTypeRead

	// Expense
	GetExpense() ([]*model.ExpenseRead, error)
	GetExpenseByID(uuid.UUID) (*model.ExpenseReadDetail, error)
}
