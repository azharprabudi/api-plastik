package query

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// SupplierQueryInterface ...
type SupplierQueryInterface interface {
	GetSuppliers(uuid.UUID) ([]*model.SupplierRead, error)
	GetSupplierByID(uuid.UUID, uuid.UUID) (*model.SupplierRead, error)
}
