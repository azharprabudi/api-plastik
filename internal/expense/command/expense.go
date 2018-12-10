package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/expense/model"
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
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

// UpdateExpense ...
func (expenseCommand *ExpenseCommand) UpdateExpense(tx *sqlx.Tx, id string, expense *model.ExpenseUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := expenseCommand.q.UpdateWhere("expenses", *expense, []*qbModel.Condition{where})

	// exec query
	_, err := tx.Exec(query, expense.ExpenseTypeID, expense.Name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpense ...
func (expenseCommand *ExpenseCommand) DeleteExpense(tx *sqlx.Tx, id string) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := expenseCommand.q.Delete("expenses", []*qbModel.Condition{where})

	// exec query
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CreateExpenseImage ...
func (expenseCommand *ExpenseCommand) CreateExpenseImage(tx *sqlx.Tx, expenseImage *model.ExpenseImageCreate) error {
	query := expenseCommand.q.Create("expense_images", *expenseImage)
	_, err := tx.Exec(query, expenseImage.Value)
	if err != nil {
		return err
	}

	return nil
}

// DeleteExpenseImage ...
func (expenseCommand *ExpenseCommand) DeleteExpenseImage(tx *sqlx.Tx, id uuid.UUID) error {
	where := &qbModel.Condition{
		Key:      "expense_id",
		Operator: "=",
		Value:    id,
		NextCond: "",
	}

	query := expenseCommand.q.Delete("expense_images", []*qbModel.Condition{where})
	_, err := tx.Exec(query)
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
