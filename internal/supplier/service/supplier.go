package service

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/transform"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/supplier/command"
	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/azharprabudi/api-plastik/internal/supplier/query"
)

// GetSupplier ...
func (ss *SupplierService) GetSupplier() ([]*dto.SupplierRes, error) {
	suppliers, err := ss.query.GetSuppliers()
	if err != nil {
		return nil, err
	}

	return ss.transform.MakeResponseGetSuppliers(suppliers), nil
}

// GetSupplierByID ...
func (ss *SupplierService) GetSupplierByID(id uuid.UUID) (*dto.SupplierRes, error) {
	supplier, err := ss.query.GetSupplierByID(id)
	if err != nil {
		return nil, err
	}

	return ss.transform.MakeResponseGetSupplierByID(supplier), nil
}

// CreateSupplier ...
func (ss *SupplierService) CreateSupplier(req *dto.SupplierReq) (uuid.UUID, error) {
	supplier := ss.transform.MakeModelCreateSupplier(req)
	err := ss.command.CreateSupplier(supplier)
	if err != nil {
		return uuid.Nil, err
	}

	return supplier.SupplierID, nil
}

// UpdateSupplier ...
func (ss *SupplierService) UpdateSupplier(id uuid.UUID, req *dto.SupplierReq) error {
	supplier := ss.transform.MakeModelUpdateSupplier(req)
	err := ss.command.UpdateSupplier(id, supplier)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSupplier ...
func (ss *SupplierService) DeleteSupplier(id uuid.UUID) error {
	err := ss.command.DeleteSupplier(id)
	if err != nil {
		return err
	}
	return nil
}

// NewSupplierService ...
func NewSupplierService(db *db.DB) SupplierServiceInterface {
	return &SupplierService{
		query:     query.NewSupplierQuery(db),
		command:   command.NewSupplierCommand(db),
		transform: transform.NewSupplierTransform(),
	}
}
