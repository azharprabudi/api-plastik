package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/item/model"
)

// GetCategory ...
func (iq *ItemQuery) GetCategory() ([]*model.ItemCategoryModelRead, error) {
	// init variable
	var results = []*model.ItemCategoryModelRead{}

	// get query
	query := iq.qb.Query("item_categories", 0, 0)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.ItemCategoryModelRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetCategoryByID ...
func (iq *ItemQuery) GetCategoryByID(categoryID int) *model.ItemCategoryModelRead {
	// init variable
	result := new(model.ItemCategoryModelRead)

	// create conditional
	where := &qbModel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    categoryID,
	}

	// get query and execute
	query := iq.qb.QueryWhere("item_categories", []*qbModel.Condition{where})
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

	// get query
	query := iq.qb.Query("items", 0, 0)
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
func (iq *ItemQuery) GetItemByID(itemID string) *model.ItemRead {
	// init variable
	result := new(model.ItemRead)

	// create conditional
	where := &qbModel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    itemID,
	}

	// get query and execute
	query := iq.qb.QueryWhere("items", []*qbModel.Condition{where})
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
