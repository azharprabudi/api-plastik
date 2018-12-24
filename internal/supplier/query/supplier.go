package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// GetSuppliers ...
func (iq *SupplierQuery) GetSuppliers(companyID uuid.UUID) ([]*model.SupplierRead, error) {
	results := []*model.SupplierRead{}
	query := iq.qb.Query("suppliers", 0, 0, []*qbmodel.Condition{
		&qbmodel.Condition{
			Key:      "company_id",
			NextCond: "",
			Operator: "=",
			Value:    companyID.String(),
		},
	}, []*qbmodel.Order{
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
func (iq *SupplierQuery) GetSupplierByID(companyID uuid.UUID, id uuid.UUID) (*model.SupplierRead, error) {
	result := new(model.SupplierRead)
	query := iq.qb.QueryWhere("suppliers", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		NextCond: "AND",
		Operator: "=",
		Value:    id.String(),
	}, &qbmodel.Condition{
		Key:      "company_id",
		NextCond: "",
		Operator: "=",
		Value:    companyID.String(),
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
