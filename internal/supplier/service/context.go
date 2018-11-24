package service

import (
	"github.com/api-plastik/internal/supplier/command"
	"github.com/api-plastik/internal/supplier/query"
)

// SupplierService ...
type SupplierService struct {
	query   query.SupplierQueryInterface
	command command.SupplierCommandInterface
}
