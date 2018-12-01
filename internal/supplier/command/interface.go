package command

import (
	"github.com/api-plastik/internal/supplier/model"
	"github.com/satori/go.uuid"
)

// SupplierCommandInterface ...
type SupplierCommandInterface interface {
	Create(*model.SupplierCreate) error
	Update(uuid.UUID, *model.SupplierUpdate) error
	Delete(uuid.UUID) error
}
