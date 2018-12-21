package transform

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
)

// MakeModelCreateExpenseType ...
func (et *ExpenseTransform) MakeModelCreateExpenseType(req *dto.ExpenseTypeReq) *model.ExpenseTypeCreate {
	return &model.ExpenseTypeCreate{
		ExpenseType: model.ExpenseType{
			ExpenseTypeID: uuid.NewV4(),
			Name:          req.Expense.Name,
			CreatedAt:     time.Now().UTC(),
		},
	}
}

// MakeModelUpdateExpenseType ...
func (et *ExpenseTransform) MakeModelUpdateExpenseType(req *dto.ExpenseTypeReq) *model.ExpenseTypeUpdate {
	return &model.ExpenseTypeUpdate{
		Name: req.Name,
	}
}

// MakeResponseGetExpenseTypes ...
func (et *ExpenseTransform) MakeResponseGetExpenseTypes(res []*model.ExpenseTypeRead) []*dto.ExpenseTypeRes {
	var results []*dto.ExpenseTypeRes
	for _, expense := range res {
		results = append(results, &dto.ExpenseTypeRes{
			ExpenseTypeID: expense.ExpenseType.ExpenseTypeID,
			Expense: dto.Expense{
				Name: expense.Name,
			},
			CreatedAt: expense.CreatedAt,
		})
	}

	return results
}

// MakeResponseGetExpenseTypeByID ...
func (et *ExpenseTransform) MakeResponseGetExpenseTypeByID(res *model.ExpenseTypeRead) *dto.ExpenseTypeRes {
	return &dto.ExpenseTypeRes{
		Expense: dto.Expense{
			Name: res.Name,
		},
		ExpenseTypeID: res.ExpenseTypeID,
		CreatedAt:     res.CreatedAt,
	}
}

// MakeModelCreateExpense ...
func (et *ExpenseTransform) MakeModelCreateExpense(req *dto.ExpenseReq) *model.ExpenseCreate {
	return &model.ExpenseCreate{
		Expense: model.Expense{
			ExpenseID:     uuid.NewV4(),
			Name:          req.Name,
			CreatedAt:     time.Now().UTC(),
			Amount:        req.Amount,
			Note:          req.Note,
			ExpenseTypeID: req.ExpenseTypeID,
		},
	}
}

// MakeResponseGetExpenses ...
func (et *ExpenseTransform) MakeResponseGetExpenses(res []*model.ExpenseRead) []*dto.ExpenseRes {
	var results []*dto.ExpenseRes
	for _, expense := range res {
		results = append(results, &dto.ExpenseRes{
			ExpenseID:     expense.Expense.ExpenseID,
			CreatedAt:     expense.Expense.CreatedAt,
			Name:          expense.Expense.Name,
			Amount:        expense.Expense.Amount,
			Note:          expense.Expense.Note,
			ExpenseTypeID: expense.Expense.ExpenseTypeID,
		})
	}

	return results
}

// MakeResponseGetExpenseByID ...
func (et *ExpenseTransform) MakeResponseGetExpenseByID(res *model.ExpenseReadDetail) *dto.ExpenseResDetail {
	var images []*dto.ExpenseImageRes
	for i := 0; i < len(res.Image); i++ {
		images = append(images, &dto.ExpenseImageRes{
			ExpenseID:      (*res.Image[i]).ExpenseID,
			ExpenseImageID: (*res.Image[i]).ExpenseImageID,
			Image:          (*res.Image[i]).Image,
		})
	}

	return &dto.ExpenseResDetail{
		ExpenseID:     res.Expense.ExpenseID,
		CreatedAt:     res.Expense.CreatedAt,
		ExpenseTypeID: res.Expense.ExpenseTypeID,
		Name:          res.Expense.Name,
		Amount:        res.Expense.Amount,
		Note:          res.Expense.Note,
		Images:        images,
	}
}

// MakeModelCreateExpenseImages ...
func (et *ExpenseTransform) MakeModelCreateExpenseImages(req *dto.ExpenseReq, id uuid.UUID) []*model.ExpenseImageCreate {
	var results []*model.ExpenseImageCreate
	for _, image := range req.Images {
		results = append(results, &model.ExpenseImageCreate{
			ExpenseImage: model.ExpenseImage{
				ExpenseID:      id,
				ExpenseImageID: uuid.NewV4(),
				Image:          image,
				CreatedAt:      time.Now().UTC(),
			},
		})
	}

	return results
}

// NewExpenseTransform ...
func NewExpenseTransform() ExpenseTransformInterface {
	return ExpenseTransformSingleton
}
