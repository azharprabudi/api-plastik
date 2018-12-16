package command

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
)

// CreateExpenseType ...
func (expenseCommand *ExpenseCommand) CreateExpenseType(expenseType *model.ExpenseTypeCreate) (int64, error) {
	// temp id for returned
	var id int64

	query := expenseCommand.q.Create("expense_types", *expenseType)
	err := expenseCommand.db.PgSQL.QueryRowx(query, expenseType.Name, expenseType.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateExpenseType ...
func (expenseCommand *ExpenseCommand) UpdateExpenseType(id int, expenseType *model.ExpenseTypeUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
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
func (expenseCommand *ExpenseCommand) DeleteExpenseType(id int) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
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
	query := expenseCommand.q.Create("expenses", *expense)
	_, err := tx.Exec(query, expense.ExpenseID, expense.ExpenseTypeID, expense.Name, expense.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// CreateExpenseImage ...
func (expenseCommand *ExpenseCommand) CreateExpenseImage(tx *sqlx.Tx, image *model.ExpenseImageCreate) error {
	query := expenseCommand.q.Create("expense_images", *image)
	_, err := tx.Exec(query, image.ExpenseImage.ExpenseImageID, image.ExpenseImage.ExepenseID, image.ExpenseImage.Image)
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
