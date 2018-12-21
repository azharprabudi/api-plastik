package service

import (
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	uuid "github.com/satori/go.uuid"
)

// ExpenseServiceInterface ...
type ExpenseServiceInterface interface {
	CreateExpenseType(*dto.ExpenseTypeReq) (uuid.UUID, error)
	UpdateExpenseType(uuid.UUID, *dto.ExpenseTypeReq) error
	DeleteExpenseType(uuid.UUID) error
	GetExpenseTypes() ([]*dto.ExpenseTypeRes, error)
	GetExpenseTypeByID(uuid.UUID) (*dto.ExpenseTypeRes, error)
	GetExpenses() ([]*dto.ExpenseRes, error)
	GetExpenseByID(uuid.UUID) (*dto.ExpenseResDetail, error)
	CreateExpense(*dto.ExpenseReq) (uuid.UUID, error)
}
