package query

import "github.com/api-plastik/internal/expense/model"

// ExpenseQueryInterface ...
type ExpenseQueryInterface interface {
	// Expense Type
	GetExpenseType() ([]*model.ExpenseTypeModelRead, error)
	GetExpenseTypeByID(int) *model.ExpenseTypeModelRead

	// Expense
	GetExpense() ([]*model.ExpenseRead, error)
	GetExpenseByID(string) *model.ExpenseRead
}
