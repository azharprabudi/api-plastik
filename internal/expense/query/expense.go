package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/expense/model"
)

// GetExpenseType ...
func (iq *ExpenseQuery) GetExpenseType() ([]*model.ExpenseTypeModelRead, error) {
	// init variable
	var results = []*model.ExpenseTypeModelRead{}

	// get query
	query := iq.qb.Query("expense_types", 0, 0)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ExpenseTypeModelRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetExpenseTypeByID ...
func (iq *ExpenseQuery) GetExpenseTypeByID(categoryID int) *model.ExpenseTypeModelRead {
	// init variable
	result := new(model.ExpenseTypeModelRead)

	// create conditional
	where := &qbModel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    categoryID,
	}

	// get query and execute
	query := iq.qb.QueryWhere("expense_types", []*qbModel.Condition{where})
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// GetExpense ...
func (iq *ExpenseQuery) GetExpense() ([]*model.ExpenseRead, error) {
	// init variable
	var results = []*model.ExpenseRead{}

	// get query
	query := iq.qb.Query("expenses", 0, 0)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ExpenseRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetExpenseByID ...
func (iq *ExpenseQuery) GetExpenseByID(expenseID string) *model.ExpenseRead {
	// init variable
	result := new(model.ExpenseRead)

	// create conditional
	where := &qbModel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    expenseID,
	}

	// get query and execute
	query := iq.qb.QueryWhere("expenses", []*qbModel.Condition{where})
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// NewExpenseQuery ...
func NewExpenseQuery(db *db.DB) ExpenseQueryInterface {
	q := qb.NewQueryBuilder()
	return &ExpenseQuery{
		qb: q,
		db: db,
	}
}
