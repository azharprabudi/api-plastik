package command

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// CreateExpenseType ...
func (expenseCommand *ExpenseCommand) CreateExpenseType(expenseType *model.ExpenseTypeCreate) error {
	query := expenseCommand.q.Create("expense_types", (*expenseType).ExpenseType)
	_, err := expenseCommand.db.PgSQL.Exec(query, expenseType.ExpenseType.ExpenseTypeID, expenseType.ExpenseType.Name, expenseType.ExpenseType.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

// UpdateExpenseType ...
func (expenseCommand *ExpenseCommand) UpdateExpenseType(id uuid.UUID, expenseType *model.ExpenseTypeUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := expenseCommand.q.UpdateWhere("expense_types", *expenseType, []*qbModel.Condition{where})

	// exec query
	_, err := expenseCommand.db.PgSQL.Exec(query, expenseType.Name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpenseType ...
func (expenseCommand *ExpenseCommand) DeleteExpenseType(id uuid.UUID) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := expenseCommand.q.Delete("expense_types", []*qbModel.Condition{where})

	// exec query
	_, err := expenseCommand.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateExpense ...
func (expenseCommand *ExpenseCommand) CreateExpense(tx *sqlx.Tx, expense *model.ExpenseCreate) error {
	query := expenseCommand.q.Create("expenses", (*expense).Expense)
	_, err := tx.Exec(query, expense.Expense.ExpenseID, expense.Expense.ExpenseTypeID, expense.Expense.Name, expense.Expense.Amount, expense.Expense.Note, expense.Expense.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// CreateExpenseImage ...
func (expenseCommand *ExpenseCommand) CreateExpenseImage(tx *sqlx.Tx, image *model.ExpenseImageCreate) error {
	query := expenseCommand.q.Create("expense_images", (*image).ExpenseImage)
	_, err := tx.Exec(query, image.ExpenseImage.ExpenseImageID, image.ExpenseImage.ExpenseID, image.ExpenseImage.Image, image.ExpenseImage.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// NewExpenseCommand ...
func NewExpenseCommand(db *db.DB) ExpenseCommandInterface {
	q := qb.NewQueryBuilder()
	return &ExpenseCommand{
		q:  q,
		db: db,
	}
}
