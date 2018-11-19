package command

import "github.com/api-plastik/db"

// NewSellerCommand ...
func NewSellerCommand(db *db.DB) SellerCommandInterface {
	return &SellerCommand{
		db: db,
	}
}
