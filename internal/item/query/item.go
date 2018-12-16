package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/satori/go.uuid"
)

// GetCategory ...
func (iq *ItemQuery) GetCategory() ([]*model.ItemCategoryRead, error) {
	// init variable
	var results = []*model.ItemCategoryRead{}

	// orders
	orders := []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	}

	// get query
	query := iq.qb.Query("item_categories", 0, 0, orders)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ItemCategoryRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetCategoryByID ...
func (iq *ItemQuery) GetCategoryByID(categoryID uuid.UUID) *model.ItemCategoryRead {
	// init variable
	result := new(model.ItemCategoryRead)

	// create conditional
	where := &qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    categoryID.String(),
	}

	// get query and execute
	query := iq.qb.QueryWhere("item_categories", []*qbmodel.Condition{where}, nil)
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// GetItem ...
func (iq *ItemQuery) GetItem() ([]*model.ItemRead, error) {
	// init variable
	var results = []*model.ItemRead{}

	// order created at
	orders := []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	}

	// get query
	query := iq.qb.Query("items", 0, 0, orders)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ItemRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetItemByID ...
func (iq *ItemQuery) GetItemByID(itemID uuid.UUID) *model.ItemRead {
	// init variable
	result := new(model.ItemRead)

	// create conditional
	where := &qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    itemID.String(),
	}

	// get query and execute
	query := iq.qb.QueryWhere("items", []*qbmodel.Condition{where}, nil)
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// NewItemQuery ...
func NewItemQuery(db *db.DB) ItemQueryInterface {
	q := qb.NewQueryBuilder()
	return &ItemQuery{
		qb: q,
		db: db,
	}
}
