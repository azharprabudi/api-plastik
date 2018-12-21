package command

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// CreateSupplier ...
func (sc *SupplierCommand) CreateSupplier(s *model.SupplierCreate) error {
	query := sc.q.Create("suppliers", (*s).Supplier)
	_, err := sc.db.PgSQL.Exec(query, s.Supplier.SupplierID, s.Supplier.Name, s.Supplier.Phone, s.Supplier.Address, s.Supplier.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSupplier ...
func (sc *SupplierCommand) UpdateSupplier(id uuid.UUID, supplier *model.SupplierUpdate) error {
	query := sc.q.UpdateWhere("suppliers", *supplier, []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}})

	_, err := sc.db.PgSQL.Exec(query, supplier.Name, supplier.Phone, supplier.Address)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSupplier ...
func (sc *SupplierCommand) DeleteSupplier(id uuid.UUID) error {
	query := sc.q.Delete("suppliers", []*qbmodel.Condition{&qbmodel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}})

	_, err := sc.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// NewSupplierCommand ...
func NewSupplierCommand(db *db.DB) SupplierCommandInterface {
	q := qb.NewQueryBuilder()
	return &SupplierCommand{
		q:  q,
		db: db,
	}
}
