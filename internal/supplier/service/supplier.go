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
func (ss *SupplierService) GetSupplier(companyID uuid.UUID) ([]*dto.SupplierRes, error) {
	suppliers, err := ss.query.GetSuppliers(companyID)
	if err != nil {
		return nil, err
	}

	return ss.transform.MakeResponseGetSuppliers(suppliers), nil
}

// GetSupplierByID ...
func (ss *SupplierService) GetSupplierByID(companyID uuid.UUID, id uuid.UUID) (*dto.SupplierRes, error) {
	supplier, err := ss.query.GetSupplierByID(companyID, id)
	if err != nil {
		return nil, err
	}

	return ss.transform.MakeResponseGetSupplierByID(supplier), nil
}

// CreateSupplier ...
func (ss *SupplierService) CreateSupplier(companyID uuid.UUID, req *dto.SupplierReq) (uuid.UUID, error) {
	supplier := ss.transform.MakeModelCreateSupplier(companyID, req)
	err := ss.command.CreateSupplier(supplier)
	if err != nil {
		return uuid.Nil, err
	}

	return supplier.SupplierID, nil
}

// UpdateSupplier ...
func (ss *SupplierService) UpdateSupplier(companyID uuid.UUID, id uuid.UUID, req *dto.SupplierReq) error {
	supplier := ss.transform.MakeModelUpdateSupplier(req)
	err := ss.command.UpdateSupplier(companyID, id, supplier)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSupplier ...
func (ss *SupplierService) DeleteSupplier(companyID uuid.UUID, id uuid.UUID) error {
	err := ss.command.DeleteSupplier(companyID, id)
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
