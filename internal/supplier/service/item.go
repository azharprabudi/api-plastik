package service

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/supplier/command"
	"github.com/api-plastik/internal/supplier/query"
)

// NewSupplierService ...
func NewSupplierService(db *db.DB) SupplierServiceInterface {
	return &SupplierService{
		query:   query.NewSupplierQuery(db),
		command: command.NewSupplierCommand(db),
	}
}
