package service

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/command"
	"github.com/api-plastik/internal/item/query"
)

// NewItemService ...
func NewItemService(db *db.DB) ItemServiceInterface {
	return &ItemService{
		query:   query.NewItemQuery(db),
		command: command.NewItemCommand(db),
	}
}
