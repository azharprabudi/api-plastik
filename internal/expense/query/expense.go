package query

import (
	"fmt"
	"time"

	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/expense/model"
	uuid "github.com/satori/go.uuid"
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
func (eq *ExpenseQuery) GetExpenseByID(expenseID uuid.UUID) (*model.ExpenseReadDetail, error) {
	// init variable
	result := new(model.ExpenseReadDetail)
	var resultImgs []*model.ExpenseImage

	// get query and execute
	rows, err := eq.db.PgSQL.Queryx(fmt.Sprintf("SELECT expenses.id as id, expenses.expense_type_id as expense_type_id, expenses.name as expense_name, expenses.amount as expense_amount, expenses.note as expense_note, expenses.created_at as expense_created_at, expense_images.id as image_id, expense_images.expense_id as expense_id, expense_images.image as image, expense_images.created_at as image_created_at FROM %s JOIN %s ON %s WHERE expenses.id = '%s'", "expenses", "expense_images", "expenses.id=expense_images.expense_id", expenseID.String()))

	if err != nil {
		return nil, err
	}

	// formatting data
	i := 0
	for rows.Next() {
		var img, expenseName, expenseNote string
		var expenseAmount float64
		var imgID, imgExpenseID, expenseID, expenseTypeID uuid.UUID
		var imgCreatedAt, expenseCreatedAt time.Time

		err := rows.Scan(&expenseID, &expenseTypeID, &expenseName, &expenseAmount, &expenseNote, &expenseCreatedAt, &imgID, &imgExpenseID, &img, &imgCreatedAt)
		fmt.Println(err)

		if i == 0 {
			result.Expense.ExpenseID = expenseID
			result.Expense.ExpenseTypeID = expenseTypeID
			result.Expense.Name = expenseName
			result.Expense.Amount = expenseAmount
			result.Expense.Note = expenseNote
			result.Expense.CreatedAt = expenseCreatedAt
		}

		resultImgs = append(resultImgs, &model.ExpenseImage{
			ExpenseID:      imgID,
			ExpenseImageID: imgExpenseID,
			Image:          img,
			CreatedAt:      imgCreatedAt,
		})
		i++
	}

	result.Image = resultImgs
	return result, nil
}

// NewExpenseQuery ...
func NewExpenseQuery(db *db.DB) ExpenseQueryInterface {
	q := qb.NewQueryBuilder()
	return &ExpenseQuery{
		qb: q,
		db: db,
	}
}
