package query

import "github.com/api-plastik/db"

// NewItemQuery ...
func NewItemQuery(db *db.DB) ItemQueryInterface {
	return &ItemQuery{
		db: db,
	}
}
