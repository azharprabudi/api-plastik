package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/item/model"
	uuid "github.com/satori/go.uuid"
)

// GetCategories ...
func (iq *ItemQuery) GetCategories(companyID uuid.UUID) ([]*model.ItemCategoryRead, error) {
	results := []*model.ItemCategoryRead{}
	query := iq.qb.Query("item_categories", 0, 0, []*qbmodel.Condition{
		&qbmodel.Condition{
			Key:      "company_id",
			NextCond: "AND",
			Operator: "=",
			Value:    companyID.String(),
		},
		&qbmodel.Condition{
			Key:      "active",
			NextCond: "",
			Operator: "=",
			Value:    true,
		},
	}, []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})

	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.ItemCategoryRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetCategoryByID ...
func (iq *ItemQuery) GetCategoryByID(companyID uuid.UUID, id uuid.UUID) (*model.ItemCategoryRead, error) {
	result := new(model.ItemCategoryRead)
	query := iq.qb.QueryWhere("item_categories", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		NextCond: "AND",
		Operator: "=",
		Value:    id.String(),
	},
		&qbmodel.Condition{
			Key:      "company_id",
			NextCond: "AND",
			Operator: "=",
			Value:    companyID.String(),
		}, &qbmodel.Condition{
			Key:      "active",
			NextCond: "",
			Operator: "=",
			Value:    true,
		}}, nil)
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetItems ...
func (iq *ItemQuery) GetItems(companyID uuid.UUID) ([]*model.ItemRead, error) {
	results := []*model.ItemRead{}
	query := iq.qb.Query("items", 0, 0, []*qbmodel.Condition{
		&qbmodel.Condition{
			Key:      "company_id",
			NextCond: "AND",
			Operator: "=",
			Value:    companyID.String(),
		},
		&qbmodel.Condition{
			Key:      "active",
			NextCond: "",
			Operator: "=",
			Value:    true,
		},
	}, []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

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
func (iq *ItemQuery) GetItemByID(companyID uuid.UUID, itemID uuid.UUID) (*model.ItemRead, error) {
	result := new(model.ItemRead)
	query := iq.qb.QueryWhere("items", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		NextCond: "AND",
		Operator: "=",
		Value:    itemID.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		NextCond: "AND",
		Operator: "=",
		Value:    companyID.String(),
	}, &qbmodel.Condition{
		Key:      "active",
		NextCond: "",
		Operator: "=",
		Value:    true,
	}}, nil)

	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetItemUnits ...
func (iq *ItemQuery) GetItemUnits() ([]*model.ItemUnitRead, error) {
	query := iq.qb.Query("item_units", 0, 0, nil, nil)
	res, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	result := []*model.ItemUnitRead{}
	for res.Next() {
		tmp := new(model.ItemUnitRead)
		res.StructScan(tmp)
		result = append(result, tmp)
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
