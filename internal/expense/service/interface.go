package service

import (
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/satori/go.uuid"
)

// ExpenseServiceInterface ...
type ExpenseServiceInterface interface {
	// Category
	CreateExpenseCategory(*dto.ExpenseCategoryReq) (int64, error)
	UpdateExpenseCategory(uuid.UUID, *dto.ExpenseCategoryReq) error
	DeleteExpenseCategory(uuid.UUID) error
	GetExpenseCategory() ([]*dto.ExpenseCategoryRes, error)
	GetExpenseCategoryByID(uuid.UUID) *dto.ExpenseCategoryRes

	// Expense
	GetExpense() ([]*dto.ExpenseRes, error)
	GetExpenseByID(uuid.UUID) *dto.ExpenseRes
	CreateExpense(*dto.ExpenseReq) (uuid.UUID, error)
}
