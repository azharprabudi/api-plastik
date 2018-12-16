package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
)

// TransformCreateExpenseType ...
func (et *ExpenseTransform) TransformCreateExpenseType(expense *dto.ExpenseTypeReq) *model.ExpenseTypeCreate {
	create := &model.ExpenseTypeCreate{
		ExpenseType: model.ExpenseType{
			Name:      expense.Name,
			CreatedAt: time.Now().UTC(),
		},
	}
	return create
}

// TransformUpdateExpenseType ...
func (et *ExpenseTransform) TransformUpdateExpenseType(expense *dto.ExpenseTypeReq) *model.ExpenseTypeUpdate {
	update := &model.ExpenseTypeUpdate{
		Name: expense.Name,
	}
	return update
}

// TransformGetExpenseType ...
func (et *ExpenseTransform) TransformGetExpenseType(expenses []*model.ExpenseTypeRead) []*dto.ExpenseTypeRes {
	// inet variable
	var res = []*dto.ExpenseTypeRes{}

	// transform data as dto expected
	for _, expense := range expenses {
		res = append(res, &dto.ExpenseTypeRes{
			ExpenseTypeID: expense.ExpenseType.ExpenseTypeID,
			ExpenseTypeReq: dto.ExpenseTypeReq{
				Name: expense.Name,
			},
			CreatedAt: expense.CreatedAt,
		})
	}

	return res
}

// TransformGetExpenseTypeByID ...
func (et *ExpenseTransform) TransformGetExpenseTypeByID(expense *model.ExpenseTypeRead) *dto.ExpenseTypeRes {
	return &dto.ExpenseTypeRes{
		ExpenseTypeReq: dto.ExpenseTypeReq{
			Name: expense.Name,
		},
		ExpenseTypeID: expense.ExpenseTypeID,
		CreatedAt:     expense.CreatedAt,
	}
}

// TransformCreateExpense ...
func (et *ExpenseTransform) TransformCreateExpense(expense *dto.ExpenseReq) *model.ExpenseCreate {
	create := &model.ExpenseCreate{
		Expense: model.Expense{
			ExpenseID:     uuid.NewV4(),
			Name:          expense.Name,
			CreatedAt:     time.Now().UTC(),
			Amount:        expense.Amount,
			Note:          expense.Note,
			ExpenseTypeID: expense.ExpenseTypeID,
			Images:        expense.Images,
		},
	}
	return create
}

// TransformGetExpense ...
func (et *ExpenseTransform) TransformGetExpense(expenses []*model.ExpenseRead) []*dto.ExpenseRes {
	// inet variable
	var res = []*dto.ExpenseRes{}

	// transform data as dto expected
	for _, expense := range expenses {
		res = append(res, &dto.ExpenseRes{
			ExpenseID: expense.Expense.ExpenseID,
			CreatedAt: expense.Expense.CreatedAt,
			ExpenseReq: dto.ExpenseReq{
				Name:          expense.Expense.Name,
				ExpenseTypeID: expense.Expense.ExpenseTypeID,
			},
		})
	}

	return res
}

// TransformGetExpenseByID ...
func (et *ExpenseTransform) TransformGetExpenseByID(expense *model.ExpenseRead) *dto.ExpenseRes {
	return &dto.ExpenseRes{
		ExpenseID: expense.Expense.ExpenseID,
		CreatedAt: expense.Expense.CreatedAt,
		ExpenseReq: dto.ExpenseReq{
			Name:          expense.Expense.Name,
			ExpenseTypeID: expense.Expense.ExpenseTypeID,
		},
	}
}

// TransformCreateExpenseImages ...
func (et *ExpenseTransform) TransformCreateExpenseImages(expenseImg []string, expenseID uuid.UUID) []*model.ExpenseImageCreate {
	var res = []*model.ExpenseImageCreate{}
	for _, image := range expenseImg {
		res = append(res, &model.ExpenseImageCreate{
			ExpenseImage: model.ExpenseImage{
				ExepenseID:     expenseID,
				ExpenseImageID: uuid.NewV4(),
				Image:          image,
			},
		})

	}
	return res
}

// NewExpenseTransform ...
func NewExpenseTransform() ExpenseTransformInterface {
	return ExpenseTransformSingleton
}
