package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
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

// NewItemQuery ...
func NewItemQuery(db *db.DB) ItemQueryInterface {
	q := qb.NewQueryBuilder()
	return &ItemQuery{
		qb: q,
		db: db,
	}
}
