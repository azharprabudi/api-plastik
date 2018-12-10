package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/api-plastik/internal/expense/dto"
	"github.com/api-plastik/internal/expense/model"
)

// TransformCreateExpenseType ...
func (et *ExpenseTransform) TransformCreateExpenseType(expenseExpenseTypeModelDTO *dto.ExpenseTypeReq) *model.ExpenseTypeCreate {
	expenseExpenseTypeCreate := &model.ExpenseTypeCreate{
		Name:      expenseExpenseTypeModelDTO.Name,
		CreatedAt: time.Now().UTC(),
	}
	return expenseExpenseTypeCreate
}

// TransformUpdateExpenseType ...
func (et *ExpenseTransform) TransformUpdateExpenseType(expenseExpenseTypeModelDTO *dto.ExpenseTypeReq) *model.ExpenseTypeUpdate {
	expenseExpenseTypeUpdate := &model.ExpenseTypeUpdate{
		Name: expenseExpenseTypeModelDTO.Name,
	}
	return expenseExpenseTypeUpdate
}

// TransformGetExpenseType ...
func (et *ExpenseTransform) TransformGetExpenseType(expenseExpenseTypeModelRead []*model.ExpenseTypeModelRead) []*dto.ExpenseTypeRes {
	// inet variable
	var expenseExpenseTypeIncRes = []*dto.ExpenseTypeRes{}

	// transform data as dto expected
	for _, expense := range expenseExpenseTypeModelRead {
		expenseExpenseTypeIncRes = append(expenseExpenseTypeIncRes, &dto.ExpenseTypeRes{
			ID:        expense.ExpenseTypeID.ExpenseTypeID,
			Name:      expense.Name,
			CreatedAt: expense.CreatedAt,
		})
	}

	return expenseExpenseTypeIncRes
}

// TransformGetExpenseTypeByID ...
func (et *ExpenseTransform) TransformGetExpenseTypeByID(expenseExpenseTypeModelRead *model.ExpenseTypeModelRead) *dto.ExpenseTypeRes {
	return &dto.ExpenseTypeRes{
		ID:        expenseExpenseTypeModelRead.ExpenseTypeID.ExpenseTypeID,
		Name:      expenseExpenseTypeModelRead.Name,
		CreatedAt: expenseExpenseTypeModelRead.CreatedAt,
	}
}

// TransformCreateExpense ...
func (et *ExpenseTransform) TransformCreateExpense(expenseDTO *dto.ExpenseReq) *model.ExpenseCreate {
	expenseCreate := &model.ExpenseCreate{
		ExpenseID:     uuid.NewV4(),
		Name:          expenseDTO.Name,
		ExpenseTypeID: expenseDTO.ExpenseTypeID,
		CreatedAt:     time.Now().UTC(),
		Amount:        expenseDTO.Amount,
		Note:          expenseDTO.Note,
	}
	return expenseCreate
}

// TransformUpdateExpense ...
func (et *ExpenseTransform) TransformUpdateExpense(expenseDTO *dto.ExpenseReq) *model.ExpenseUpdate {
	expenseUpdate := &model.ExpenseUpdate{
		Name:          expenseDTO.Name,
		ExpenseTypeID: expenseDTO.ExpenseTypeID,
		Amount:        expenseDTO.Amount,
		Note:          expenseDTO.Note,
	}
	return expenseUpdate
}

// TransformGetExpense ...
func (et *ExpenseTransform) TransformGetExpense(expenseRead []*model.ExpenseRead) []*dto.ExpenseRes {
	// inet variable
	var expenseRes = []*dto.ExpenseRes{}

	// transform data as dto expected
	for _, expense := range expenseRead {
		expenseRes = append(expenseRes, &dto.ExpenseRes{
			ID:        expense.ExpenseCreate.ExpenseID,
			CreatedAt: expense.ExpenseCreate.CreatedAt,
			ExpenseReq: dto.ExpenseReq{
				Name:          expense.ExpenseCreate.Name,
				ExpenseTypeID: expense.ExpenseCreate.ExpenseTypeID,
			},
		})
	}

	return expenseRes
}

// TransformGetExpenseByID ...
func (et *ExpenseTransform) TransformGetExpenseByID(expenseRead *model.ExpenseRead) *dto.ExpenseRes {
	return &dto.ExpenseRes{
		ID:        expenseRead.ExpenseCreate.ExpenseID,
		CreatedAt: expenseRead.ExpenseCreate.CreatedAt,
		ExpenseReq: dto.ExpenseReq{
			Name:          expenseRead.ExpenseCreate.Name,
			ExpenseTypeID: expenseRead.ExpenseCreate.ExpenseTypeID,
		},
	}
}

// NewExpenseTransform ...
func NewExpenseTransform() ExpenseTransformInterface {
	return ExpenseTransformSingleton
}
