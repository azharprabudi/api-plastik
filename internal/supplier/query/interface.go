package query

import (
	"github.com/api-plastik/internal/supplier/model"
	"github.com/satori/go.uuid"
)

// SupplierQueryInterface ...
type SupplierQueryInterface interface {
	Get() ([]*model.SupplierRead, error)
	GetByID(uuid.UUID) *model.SupplierRead
}
