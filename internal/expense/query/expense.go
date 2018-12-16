package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	"github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	"github.com/satori/go.uuid"
)

// GetExpenseType ...
func (eq *ExpenseQuery) GetExpenseType() ([]*model.ExpenseTypeRead, error) {
	// init variable
	var results = []*model.ExpenseTypeRead{}

	// order
	orders := []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	}

	// get query
	query := eq.qb.Query("expense_types", 0, 0, orders)
	rows, err := eq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ExpenseTypeRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetExpenseTypeByID ...
func (eq *ExpenseQuery) GetExpenseTypeByID(ExpenseTypeID uuid.UUID) *model.ExpenseTypeRead {
	// init variable
	result := new(model.ExpenseTypeRead)

	// create conditional
	where := &qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    ExpenseTypeID.String(),
	}

	// get query and execute
	query := eq.qb.QueryWhere("expense_types", []*qbmodel.Condition{where}, nil)
	err := eq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// GetExpense ...
func (eq *ExpenseQuery) GetExpense() ([]*model.ExpenseRead, error) {
	// init variable
	var results = []*model.ExpenseRead{}

	// orders
	orders := []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	}

	// get query
	query := eq.qb.Query("expenses", 0, 0, orders)
	rows, err := eq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ExpenseRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetExpenseByID ...
func (eq *ExpenseQuery) GetExpenseByID(expenseID uuid.UUID) *model.ExpenseReadDetail {
	// init variable
	result := new(model.ExpenseReadDetail)

	// create conditional
	where := &qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    expenseID,
	}

	// create join
	join := &qbmodel.Join{
		TableFrom:       "",
		ColumnTableFrom: "id",
		TableWith:       "expense_images",
		ColumnTableWith: "expense_id",
	}

	// get query and execute
	query := eq.qb.QueryWhereWith("expenses", []*qbmodel.Join{join}, []*qbmodel.Condition{where}, nil)
	err := eq.db.PgSQL.QueryRowx(query).StructScan(result)
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
