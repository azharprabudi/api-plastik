package service

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/command"
	"github.com/azharprabudi/api-plastik/internal/supplier/query"
	"github.com/azharprabudi/api-plastik/internal/supplier/transform"
)

// SupplierService ...
type SupplierService struct {
	query     query.SupplierQueryInterface
	command   command.SupplierCommandInterface
	transform transform.SupplierTransformInterface
}
