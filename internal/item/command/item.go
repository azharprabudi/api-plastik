package command

import "github.com/api-plastik/db"

// NewItemCommand ...
func NewItemCommand(db *db.DB) ItemCommandInterface {
	return &ItemCommand{
		db: db,
	}
}
