package service

import (
	"time"

	"github.com/api-plastik/internal/item/transform"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/command"
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
	"github.com/api-plastik/internal/item/query"
)

// GetItemCategory ...
func (itemService *ItemService) GetItemCategory() ([]*dto.ItemCategoryIncRes, error) {
	// add data to db
	categoriesModel, err := itemService.query.GetCategory()
	if err != nil {
		return nil, err
	}

	// transform data from model
	categoriesDTO := itemService.transform.TransformGetCategory(categoriesModel)
	return categoriesDTO, nil
}

// GetItemCategoryByID ...
func (itemService *ItemService) GetItemCategoryByID(categoryID int) (*dto.ItemCategoryIncRes, error) {
	categoryModel, err := itemService.query.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	// transform data from model
	categoryDTO := itemService.transform.TransformGetCategoryByID(categoryModel)
	return categoryDTO, nil
}

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
		query:     query.NewItemQuery(db),
		command:   command.NewItemCommand(db),
		transform: transform.NewItemTransform(),
	}
}
