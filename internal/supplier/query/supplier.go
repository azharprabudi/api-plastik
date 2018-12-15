package query

import (
	"github.com/api-plastik/db"
	qb "github.com/api-plastik/helper/querybuilder"
	"github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// Get ...
func (iq *SupplierQuery) Get() ([]*model.SupplierRead, error) {
	// init variable
	var results = []*model.SupplierRead{}

	// ordering data
	orders := []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "DESC",
		},
	}

	// get query
	query := iq.qb.Query("suppliers", 0, 0, orders)
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	// get struct
	for rows.Next() {
		tmp := new(model.SupplierRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}
	if err != nil {
		return nil, err
	}

	return results, nil
}

// GetByID ...
func (iq *SupplierQuery) GetByID(supplierID uuid.UUID) *model.SupplierRead {
	// init variable
	result := new(model.SupplierRead)

	// create conditional
	where := &qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    supplierID.String(),
	}

	// get query and execute
	query := iq.qb.QueryWhere("suppliers", []*qbmodel.Condition{where}, nil)
	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil
	}
	return result
}

// NewSupplierQuery ...
func NewSupplierQuery(db *db.DB) SupplierQueryInterface {
	q := qb.NewQueryBuilder()
	return &SupplierQuery{
		qb: q,
		db: db,
	}
}
