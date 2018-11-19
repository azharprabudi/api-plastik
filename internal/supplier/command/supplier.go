package command

import "github.com/api-plastik/db"

// NewSupplierCommand ...
func NewSupplierCommand(db *db.DB) SupplierCommandInterface {
	return &SupplierCommand{
		db: db,
	}
}
