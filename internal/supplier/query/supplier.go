package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// GetSuppliers ...
func (iq *SupplierQuery) GetSuppliers() ([]*model.SupplierRead, error) {
	var results = []*model.SupplierRead{}
	query := iq.qb.Query("suppliers", 0, 0, []*qbmodel.Order{
		&qbmodel.Order{
			Key:   "created_at",
			Value: "DESC",
		},
	})
	rows, err := iq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

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

// GetSupplierByID ...
func (iq *SupplierQuery) GetSupplierByID(id uuid.UUID) (*model.SupplierRead, error) {
	result := new(model.SupplierRead)
	query := iq.qb.QueryWhere("suppliers", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    id.String(),
	}}, nil)

	err := iq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewSupplierQuery ...
func NewSupplierQuery(db *db.DB) SupplierQueryInterface {
	q := qb.NewQueryBuilder()
	return &SupplierQuery{
		qb: q,
		db: db,
	}
}
