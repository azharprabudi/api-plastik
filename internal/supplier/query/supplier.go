package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/supplier/model"
	"github.com/satori/go.uuid"
)

// Get ...
func (iq *SupplierQuery) Get() ([]*model.SupplierRead, error) {
	// init variable
	var results = []*model.SupplierRead{}

	// get query
	query := iq.qb.Query("suppliers", 0, 0)
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
	where := &qbModel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    supplierID,
	}

	// get query and execute
	query := iq.qb.QueryWhere("suppliers", []*qbModel.Condition{where})
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
