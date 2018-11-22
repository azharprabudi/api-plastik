package service

import (
	"time"

	"github.com/api-plastik/db"
	"github.com/api-plastik/dto"
	"github.com/api-plastik/internal/item/command"
	"github.com/api-plastik/internal/item/model"
	"github.com/api-plastik/internal/item/query"
)

// CreateItemCategory ...
func (itemService *ItemService) CreateItemCategory(itemCategory *dto.ItemCategoryIncReq) error {
	itemCategoryCreate := &model.ItemCategoryCreate{
		Name:      itemCategory.Name,
		CreatedAt: time.Now().UTC(),
	}

	// add data to db
	err := itemService.command.CreateCategory(itemCategoryCreate)
	if err != nil {
		return err
	}
	return nil
}

// NewItemService ...
func NewItemService(db *db.DB) ItemServiceInterface {
	return &ItemService{
		query:   query.NewItemQuery(db),
		command: command.NewItemCommand(db),
	}
}
