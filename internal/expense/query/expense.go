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

// GetExpenseTypes ...
func (eq *ExpenseQuery) GetExpenseTypes() ([]*model.ExpenseTypeRead, error) {
	var results []*model.ExpenseTypeRead
	query := eq.qb.Query("expense_types", 0, 0, []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})
	rows, err := eq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.ExpenseTypeRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetExpenseTypeByID ...
func (eq *ExpenseQuery) GetExpenseTypeByID(id uuid.UUID) (*model.ExpenseTypeRead, error) {
	result := new(model.ExpenseTypeRead)
	query := eq.qb.QueryWhere("expense_types", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    id.String(),
	}}, nil)
	err := eq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetExpenses ...
func (eq *ExpenseQuery) GetExpenses() ([]*model.ExpenseRead, error) {
	var results []*model.ExpenseRead
	query := eq.qb.Query("expenses", 0, 0, []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})
	rows, err := eq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.ExpenseRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetExpenseByID ...
func (eq *ExpenseQuery) GetExpenseByID(id uuid.UUID) (*model.ExpenseReadDetail, error) {
	result := new(model.ExpenseReadDetail)
	var resultImages []*model.ExpenseImage

	rows, err := eq.db.PgSQL.Queryx(fmt.Sprintf("SELECT expenses.id as id, expenses.expense_type_id as expense_type_id, expenses.name as expense_name, expenses.amount as expense_amount, expenses.note as expense_note, expenses.created_at as expense_created_at, expense_images.id as image_id, expense_images.expense_id as expense_id, expense_images.image as image, expense_images.created_at as image_created_at FROM %s JOIN %s ON %s WHERE expenses.id = '%s'", "expenses", "expense_images", "expenses.id=expense_images.expense_id", id.String()))
	if err != nil {
		return nil, err
	}

	i := 0
	for rows.Next() {
		var amount float64
		var image, name, note string
		var imageCreated, expenseCreated time.Time
		var imageID, imageExpenseID, expenseID, expenseTypeID uuid.UUID

		rows.Scan(&expenseID, &expenseTypeID, &name, &amount, &note, &expenseCreated, &imageID, &imageExpenseID, &image, &imageCreated)

		if i == 0 {
			result.Expense.ExpenseID = expenseID
			result.Expense.ExpenseTypeID = expenseTypeID
			result.Expense.Name = name
			result.Expense.Amount = amount
			result.Expense.Note = note
			result.Expense.CreatedAt = expenseCreated
		}

		resultImages = append(resultImages, &model.ExpenseImage{
			ExpenseID:      expenseID,
			ExpenseImageID: imageID,
			Image:          image,
			CreatedAt:      imageCreated,
		})
		i++
	}

	result.Image = resultImages
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
