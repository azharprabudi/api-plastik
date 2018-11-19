package query

import "github.com/api-plastik/db"

// NewSupplierQuery ...
func NewSupplierQuery(db *db.DB) SupplierQueryInterface {
	return &SupplierQuery{
		db: db,
	}
}
