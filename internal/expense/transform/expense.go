package transform

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
)

// TransformCreateExpenseType ...
func (et *ExpenseTransform) TransformCreateExpenseType(expense *dto.ExpenseTypeReq) *model.ExpenseTypeCreate {
	create := &model.ExpenseTypeCreate{
		ExpenseType: model.ExpenseType{
			ExpenseTypeID: uuid.NewV4(),
			Name:          expense.Name,
			CreatedAt:     time.Now().UTC(),
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
			ExpenseID:     expense.Expense.ExpenseID,
			CreatedAt:     expense.Expense.CreatedAt,
			Name:          expense.Expense.Name,
			Amount:        expense.Expense.Amount,
			Note:          expense.Expense.Note,
			ExpenseTypeID: expense.Expense.ExpenseTypeID,
		})
	}

	return res
}

// TransformGetExpenseByID ...
func (et *ExpenseTransform) TransformGetExpenseByID(expense *model.ExpenseReadDetail) *dto.ExpenseResDetail {
	var images = []*dto.ExpenseImageRes{}

	// loop the model images
	for i := 0; i < len(expense.Image); i++ {
		images = append(images, &dto.ExpenseImageRes{
			ExpenseID:      (*expense.Image[i]).ExpenseID,
			ExpenseImageID: (*expense.Image[i]).ExpenseImageID,
			Image:          (*expense.Image[i]).Image,
		})
	}

	return &dto.ExpenseResDetail{
		ExpenseID:     expense.Expense.ExpenseID,
		CreatedAt:     expense.Expense.CreatedAt,
		ExpenseTypeID: expense.Expense.ExpenseTypeID,
		Name:          expense.Expense.Name,
		Amount:        expense.Expense.Amount,
		Note:          expense.Expense.Note,
		Images:        images,
	}
}

// TransformCreateExpenseImages ...
func (et *ExpenseTransform) TransformCreateExpenseImages(expenseImg []string, expenseID uuid.UUID) []*model.ExpenseImageCreate {
	var res = []*model.ExpenseImageCreate{}
	for _, image := range expenseImg {
		res = append(res, &model.ExpenseImageCreate{
			ExpenseImage: model.ExpenseImage{
				ExpenseID:      expenseID,
				ExpenseImageID: uuid.NewV4(),
				Image:          image,
				CreatedAt:      time.Now().UTC(),
			},
		})
	}

	return res
}

// NewExpenseTransform ...
func NewExpenseTransform() ExpenseTransformInterface {
	return ExpenseTransformSingleton
}
