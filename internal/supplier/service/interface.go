package service

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	uuid "github.com/satori/go.uuid"
)

// SupplierServiceInterface ...
type SupplierServiceInterface interface {
	CreateSupplier(uuid.UUID, *dto.SupplierReq) (uuid.UUID, error)
	UpdateSupplier(uuid.UUID, uuid.UUID, *dto.SupplierReq) error
	DeleteSupplier(uuid.UUID, uuid.UUID) error
	GetSupplier(uuid.UUID) ([]*dto.SupplierRes, error)
	GetSupplierByID(uuid.UUID, uuid.UUID) (*dto.SupplierRes, error)
}
