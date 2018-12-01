package service

import (
	"github.com/api-plastik/internal/supplier/command"
	"github.com/api-plastik/internal/supplier/query"
	"github.com/api-plastik/internal/supplier/transform"
)

// SupplierService ...
type SupplierService struct {
	query     query.SupplierQueryInterface
	command   command.SupplierCommandInterface
	transform transform.SupplierTransformInterface
}
