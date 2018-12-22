package service

import (
	trx "github.com/azharprabudi/api-plastik/helper/transaction"
	"github.com/azharprabudi/api-plastik/internal/expense/transform"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/expense/command"
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/query"
)

// GetExpenseType ...
func (es *ExpenseService) GetExpenseTypes() ([]*dto.ExpenseTypeRes, error) {
	expenseTypes, err := es.query.GetExpenseTypes()
	if err != nil {
		return nil, err
	}

	return es.transform.MakeResponseGetExpenseTypes(expenseTypes), nil
}

// GetExpenseTypeByID ...
func (es *ExpenseService) GetExpenseTypeByID(id uuid.UUID) (*dto.ExpenseTypeRes, error) {
	expense, err := es.query.GetExpenseTypeByID(id)
	if err != nil {
		return nil, err
	}

	return es.transform.MakeResponseGetExpenseTypeByID(expense), nil
}

// CreateExpenseType ...
func (es *ExpenseService) CreateExpenseType(req *dto.ExpenseTypeReq) (uuid.UUID, error) {
	expenseType := es.transform.MakeModelCreateExpenseType(req)
	err := es.command.CreateExpenseType(expenseType)
	if err != nil {
		return uuid.Nil, err
	}

	return expenseType.ExpenseType.ExpenseTypeID, nil
}

// UpdateExpenseType ...
func (es *ExpenseService) UpdateExpenseType(id uuid.UUID, req *dto.ExpenseTypeReq) error {
	expenseType := es.transform.MakeModelUpdateExpenseType(req)
	err := es.command.UpdateExpenseType(id, expenseType)
	if err != nil {
		return err
	}

	return nil
}

// DeleteExpenseType ...
func (es *ExpenseService) DeleteExpenseType(id uuid.UUID) error {
	err := es.command.DeleteExpenseType(id)
	if err != nil {
		return err
	}
	return nil
}

// GetExpenses ...
func (es *ExpenseService) GetExpenses() ([]*dto.ExpenseRes, error) {
	expense, err := es.query.GetExpenses()
	if err != nil {
		return nil, err
	}

	return es.transform.MakeResponseGetExpenses(expense), nil
}

// GetExpenseByID ...
func (es *ExpenseService) GetExpenseByID(id uuid.UUID) (*dto.ExpenseResDetail, error) {
	expense, err := es.query.GetExpenseByID(id)
	if err != nil {
		return nil, err
	}

	return es.transform.MakeResponseGetExpenseByID(expense), nil
}

// CreateExpense ...
func (es *ExpenseService) CreateExpense(req *dto.ExpenseReq) (uuid.UUID, error) {
	expense := es.transform.MakeModelCreateExpense(req)
	expenseImages := es.transform.MakeModelCreateExpenseImages(req, expense.Expense.ExpenseID)

	newTrx := trx.NewTransaction()
	ctx := newTrx.CreateTrx(es.db.PgSQL)
	err := newTrx.RunTrx(ctx, func(tx *sqlx.Tx) error {
		err := es.command.CreateExpense(tx, expense)
		if err != nil {
			return err
		}

		for _, image := range expenseImages {
			err = es.command.CreateExpenseImage(tx, image)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}
	return expense.Expense.ExpenseID, nil
}

// NewExpenseService ...
func NewExpenseService(db *db.DB) ExpenseServiceInterface {
	return &ExpenseService{
		db:        db,
		query:     query.NewExpenseQuery(db),
		command:   command.NewExpenseCommand(db),
		transform: transform.NewExpenseTransform(),
	}
}
