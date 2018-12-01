package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	"github.com/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// Create ...
func (supplierCommand *SupplierCommand) Create(supplier *model.SupplierCreate) error {
	query := supplierCommand.q.Create("suppliers", *supplier)
	_, err := supplierCommand.db.PgSQL.Exec(query, supplier.Name, supplier.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (supplierCommand *SupplierCommand) Update(id uuid.UUID, supplier *model.SupplierUpdate) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := supplierCommand.q.UpdateWhere("suppliers", *supplier, []*qbModel.Condition{where})

	// exec query
	_, err := supplierCommand.db.PgSQL.Exec(query, supplier.Name)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (supplierCommand *SupplierCommand) Delete(id uuid.UUID) error {
	// create condition
	where := &qbModel.Condition{
		Key:      "id",
		Value:    id,
		Operator: "=",
		NextCond: "",
	}

	// create query
	query := supplierCommand.q.Delete("suppliers", []*qbModel.Condition{where})

	// exec query
	_, err := supplierCommand.db.PgSQL.Exec(query)
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
