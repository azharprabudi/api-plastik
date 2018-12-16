package command

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// Create ...
func (sc *SupplierCommand) Create(s *model.SupplierCreate) error {
	query := sc.q.Create("suppliers", (*s).Supplier)
	_, err := sc.db.PgSQL.Exec(query, s.Supplier.SupplierID, s.Supplier.Name, s.Supplier.Phone, s.Supplier.Address, s.Supplier.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (sc *SupplierCommand) Update(id uuid.UUID, supplier *model.SupplierUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := sc.q.UpdateWhere("suppliers", *supplier, []*qbModel.Condition{where})

	// exec query
	_, err := sc.db.PgSQL.Exec(query, supplier.Name, supplier.Phone, supplier.Address)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (sc *SupplierCommand) Delete(id uuid.UUID) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "",
		Value:    id.String(),
	}

	// create query
	query := sc.q.Delete("suppliers", []*qbModel.Condition{where})

	// exec query
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
