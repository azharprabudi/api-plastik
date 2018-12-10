package service

import (
	"github.com/api-plastik/internal/expense/dto"
	"github.com/satori/go.uuid"
)

// ExpenseServiceInterface ...
type ExpenseServiceInterface interface {
	// Category
	CreateExpenseCategory(*dto.ExpenseCategoryReq) (int64, error)
	UpdateExpenseCategory(int, *dto.ExpenseCategoryReq) error
	DeleteExpenseCategory(int) error
	GetExpenseCategory() ([]*dto.ExpenseCategoryRes, error)
	GetExpenseCategoryByID(int) *dto.ExpenseCategoryRes

	// Expense
	CreateExpense(*dto.ExpenseReq) (uuid.UUID, error)
	UpdateExpense(string, *dto.ExpenseReq) error
	DeleteExpense(string) error
	GetExpense() ([]*dto.ExpenseRes, error)
	GetExpenseByID(string) *dto.ExpenseRes
}
