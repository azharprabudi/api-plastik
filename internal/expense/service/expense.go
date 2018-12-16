package service

import (
	"github.com/azharprabudi/api-plastik/helper/transaction"
	"github.com/azharprabudi/api-plastik/internal/expense/transform"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/expense/command"
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/query"
)

// GetExpenseType ...
func (es *ExpenseService) GetExpenseType() ([]*dto.ExpenseTypeRes, error) {
	// add data to db
	expenseType, err := es.query.GetExpenseType()
	if err != nil {
		return nil, err
	}

	// transform data from model
	expenseTypeDTO := es.transform.TransformGetExpenseType(expenseType)
	return expenseTypeDTO, nil
}

// GetExpenseTypeByID ...
func (es *ExpenseService) GetExpenseTypeByID(expenseTypeID uuid.UUID) *dto.ExpenseTypeRes {
	expenseType := es.query.GetExpenseTypeByID(expenseTypeID)
	if expenseType == nil {
		return nil
	}

	// transform data from model
	expenseTypeDTO := es.transform.TransformGetExpenseTypeByID(expenseType)
	return expenseTypeDTO
}

// CreateExpenseType ...
func (es *ExpenseService) CreateExpenseType(itemCategory *dto.ExpenseTypeReq) (uuid.UUID, error) {
	create := es.transform.TransformCreateCategory(itemCategory)

	// add data to db
	id, err := es.command.CreateCategory(create)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// UpdateExpenseType ...
func (es *ExpenseService) UpdateExpenseType(expenseTypeID uuid.UUID, itemCategory *dto.ExpenseTypeReq) error {
	update := es.transform.TransformUpdateCategory(itemCategory)

	// update to db
	err := es.command.UpdateCategory(expenseTypeID, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpenseType ...
func (es *ExpenseService) DeleteExpenseType(expenseTypeID uuid.UUID) error {
	// delete data from db
	err := es.command.DeleteCategory(expenseTypeID)
	if err != nil {
		return err
	}
	return nil
}

// GetExpense ...
func (es *ExpenseService) GetExpense() ([]*dto.ExpenseRes, error) {
	// add data to db
	expense, err := es.query.GetExpense()
	if err != nil {
		return nil, err
	}

	// transform data from model
	itemDTO := es.transform.TransformGetExpense(expense)
	return itemDTO, nil
}

// GetExpenseByID ...
func (es *ExpenseService) GetExpenseByID(expenseID uuid.UUID) *dto.ExpenseRes {
	expense := es.query.GetExpenseByID(expenseID)
	if expense == nil {
		return nil
	}

	// transform data from model
	itemDTO := es.transform.TransformGetExpenseByID(expense)
	return itemDTO
}

// CreateExpense ...
func (es *ExpenseService) CreateExpense(expense *dto.ExpenseReq) (uuid.UUID, error) {
	// create expense
	create := es.transform.TransformCreateExpense(expense)

	// create expense images
	createImg := es.transform.TransformCreateExpenseImages(expense.Images)

	// add data to db
	newTrx := trx.NewTransaction()
	ctx := newTrx.CreateTrx()
	newTrx.RunTrx(ctx, func(tx *sqlx.Tx) {
		// insert expense
		err := es.command.CreateExpense(tx, create)
		if err != nil {
			return err
		}

		// insert images
		for _, expenseImg := range createImg {
			err := es.command.CreateExpenseImage(tx, expenseImg)
			if err != nil {
				break
			}
		}

		// return error if exists
		if err != nil {
			return err
		}
		return nil

	})

	return create.ExpenseID, nil
}

// NewExpenseService ...
func NewExpenseService(db *db.DB) ExpenseServiceInterface {
	return &ExpenseService{
		query:     query.NewExpenseQuery(db),
		command:   command.NewExpenseCommand(db),
		transform: transform.NewExpenseTransform(),
	}
}
