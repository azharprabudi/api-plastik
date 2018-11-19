package query

import "github.com/api-plastik/db"

// NewSellerQuery ...
func NewSellerQuery(db *db.DB) SellerQueryInterface {
	return &SellerQuery{
		db: db,
	}
}
