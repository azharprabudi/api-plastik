package query

import (
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	uuid "github.com/satori/go.uuid"
)

// ExpenseQueryInterface ...
type ExpenseQueryInterface interface {
	GetExpenseTypes() ([]*model.ExpenseTypeRead, error)
	GetExpenseTypeByID(uuid.UUID) (*model.ExpenseTypeRead, error)
	GetExpenses() ([]*model.ExpenseRead, error)
	GetExpenseByID(uuid.UUID) (*model.ExpenseReadDetail, error)
}
