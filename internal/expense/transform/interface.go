package transform

import (
	"github.com/api-plastik/internal/expense/dto"
	"github.com/api-plastik/internal/expense/model"
)

// ExpenseTransformInterface ...
type ExpenseTransformInterface interface {
	// category
	TransformCreateExpenseType(*dto.ExpenseTypeReq) *model.ExpenseTypeCreate
	TransformUpdateExpenseType(*dto.ExpenseTypeReq) *model.ExpenseTypeUpdate
	TransformGetExpenseType([]*model.ExpenseTypeModelRead) []*dto.ExpenseTypeRes
	TransformGetExpenseTypeByID(*model.ExpenseTypeModelRead) *dto.ExpenseTypeRes

	// expense
	TransformCreateExpense(*dto.ExpenseReq) *model.ExpenseCreate
	TransformUpdateExpense(*dto.ExpenseReq) *model.ExpenseUpdate
	TransformGetExpense([]*model.ExpenseRead) []*dto.ExpenseRes
	TransformGetExpenseByID(*model.ExpenseRead) *dto.ExpenseRes
}
