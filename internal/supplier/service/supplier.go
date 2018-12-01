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
func (supplierService *SupplierService) GetSupplier() ([]*dto.SupplierRes, error) {
	// add data to db
	supplierModel, err := supplierService.query.Get()
	if err != nil {
		return nil, err
	}

	// transform data from model
	supplierDTO := supplierService.transform.TransformGet(supplierModel)
	return supplierDTO, nil
}

// GetSupplierByID ...
func (supplierService *SupplierService) GetSupplierByID(supplierID uuid.UUID) *dto.SupplierRes {
	supplierModel := supplierService.query.GetByID(supplierID)
	if supplierModel == nil {
		return nil
	}

	// transform data from model
	supplierDTO := supplierService.transform.TransformGetByID(supplierModel)
	return supplierDTO
}

// CreateSupplier ...
func (supplierService *SupplierService) CreateSupplier(item *dto.SupplierReq) (uuid.UUID, error) {
	// transform dto to model
	supplierCreate := supplierService.transform.TransformCreate(item)

	// add data to db
	err := supplierService.command.Create(supplierCreate)
	if err != nil {
		return uuid.Nil, err
	}
	return supplierCreate.SupplierID, nil
}

// UpdateSupplier ...
func (supplierService *SupplierService) UpdateSupplier(supplierID uuid.UUID, item *dto.SupplierReq) error {
	// transform dto to model
	supplierUpdate := supplierService.transform.TransformUpdate(item)

	// update to db
	err := supplierService.command.Update(supplierID, supplierUpdate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSupplier ...
func (supplierService *SupplierService) DeleteSupplier(supplierID uuid.UUID) error {
	// delete data from db
	err := supplierService.command.Delete(supplierID)
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
