package service

import (
	"github.com/api-plastik/internal/supplier/transform"
	"github.com/satori/go.uuid"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/supplier/command"
	"github.com/api-plastik/internal/supplier/dto"
	"github.com/api-plastik/internal/supplier/query"
)

// GetSupplier ...
func (ss *SupplierService) GetSupplier() ([]*dto.SupplierRes, error) {
	// add data to db
	supplier, err := ss.query.Get()
	if err != nil {
		return nil, err
	}

	// transform data from model
	supplierDTO := ss.transform.TransformGet(supplier)
	return supplierDTO, nil
}

// GetSupplierByID ...
func (ss *SupplierService) GetSupplierByID(supplierID uuid.UUID) *dto.SupplierRes {
	supplier := ss.query.GetByID(supplierID)
	if supplier == nil {
		return nil
	}

	// transform data from model
	supplierDTO := ss.transform.TransformGetByID(supplier)
	return supplierDTO
}

// CreateSupplier ...
func (ss *SupplierService) CreateSupplier(item *dto.SupplierReq) (uuid.UUID, error) {
	// transform dto to model
	supplier := ss.transform.TransformCreate(item)

	// add data to db
	err := ss.command.Create(supplier)
	if err != nil {
		return uuid.Nil, err
	}
	return supplier.SupplierID, nil
}

// UpdateSupplier ...
func (ss *SupplierService) UpdateSupplier(supplierID uuid.UUID, item *dto.SupplierReq) error {
	// transform dto to model
	supplier := ss.transform.TransformUpdate(item)

	// update to db
	err := ss.command.Update(supplierID, supplier)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSupplier ...
func (ss *SupplierService) DeleteSupplier(supplierID uuid.UUID) error {
	// delete data from db
	err := ss.command.Delete(supplierID)
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
