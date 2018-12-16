package service

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/satori/go.uuid"
)

// SupplierServiceInterface ...
type SupplierServiceInterface interface {
	CreateSupplier(*dto.SupplierReq) (uuid.UUID, error)
	UpdateSupplier(uuid.UUID, *dto.SupplierReq) error
	DeleteSupplier(uuid.UUID) error
	GetSupplier() ([]*dto.SupplierRes, error)
	GetSupplierByID(uuid.UUID) *dto.SupplierRes
}
