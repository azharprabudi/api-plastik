package transform

import (
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/satori/go.uuid"
)

// ExpenseTransformInterface ...
type ExpenseTransformInterface interface {
	// expense type
	TransformGetExpenseType([]*model.ExpenseTypeRead) []*dto.ExpenseTypeRes
	TransformGetExpenseTypeByID(*model.ExpenseTypeRead) *dto.ExpenseTypeRes
	TransformCreateExpenseType(*dto.ExpenseTypeReq) *model.ExpenseTypeCreate
	TransformUpdateExpenseType(*dto.ExpenseTypeReq) *model.ExpenseTypeUpdate

	// expense
	TransformCreateExpense(*dto.ExpenseReq) *model.ExpenseCreate
	TransformGetExpense([]*model.ExpenseRead) []*dto.ExpenseRes
	TransformGetExpenseByID(*model.ExpenseRead) *dto.ExpenseRes

	// expense images
	TransformCreateExpenseImages([]string, uuid.UUID) []*model.ExpenseImageCreate
}
