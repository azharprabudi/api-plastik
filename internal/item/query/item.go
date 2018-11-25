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
	query := iq.qb.Query("itemCategory", 0, 0)
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
func (iq *ItemQuery) GetCategoryByID(categoryID int) (*model.ItemCategoryModelRead, error) {
	// init variable
	result := new(model.ItemCategoryModelRead)

	// create conditional
	where := &qbModel.Condition{
		Key:      "itemCategoryId",
		NextCond: "",
		Operator: "=",
		Value:    categoryID,
	}

	// get query and execute
	query := iq.qb.QueryWhere("itemCategory", []*qbModel.Condition{where})
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewItemQuery ...
func NewItemQuery(db *db.DB) ItemQueryInterface {
	q := qb.NewQueryBuilder()
	return &ItemQuery{
		qb: q,
		db: db,
	}
}
