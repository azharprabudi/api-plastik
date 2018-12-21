package transform

import (
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	uuid "github.com/satori/go.uuid"
)

// ExpenseTransformInterface ...
type ExpenseTransformInterface interface {
	MakeResponseGetExpenseTypes([]*model.ExpenseTypeRead) []*dto.ExpenseTypeRes
	MakeResponseGetExpenseTypeByID(*model.ExpenseTypeRead) *dto.ExpenseTypeRes
	MakeModelCreateExpenseType(*dto.ExpenseTypeReq) *model.ExpenseTypeCreate
	MakeModelUpdateExpenseType(*dto.ExpenseTypeReq) *model.ExpenseTypeUpdate
	MakeModelCreateExpense(*dto.ExpenseReq) *model.ExpenseCreate
	MakeResponseGetExpenses([]*model.ExpenseRead) []*dto.ExpenseRes
	MakeResponseGetExpenseByID(*model.ExpenseReadDetail) *dto.ExpenseResDetail
	MakeModelCreateExpenseImages(*dto.ExpenseReq, uuid.UUID) []*model.ExpenseImageCreate
}
