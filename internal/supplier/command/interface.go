package command

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// SupplierCommandInterface ...
type SupplierCommandInterface interface {
	CreateSupplier(*model.SupplierCreate) error
	UpdateSupplier(uuid.UUID, uuid.UUID, *model.SupplierUpdate) error
	DeleteSupplier(uuid.UUID, uuid.UUID) error
}
