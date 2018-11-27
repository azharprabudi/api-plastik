package service

import (
	"github.com/api-plastik/internal/item/transform"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/command"
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/query"
)

// GetItemCategory ...
func (itemService *ItemService) GetItemCategory() ([]*dto.ItemCategoryRes, error) {
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
func (itemService *ItemService) GetItemCategoryByID(categoryID int) *dto.ItemCategoryRes {
	categoryModel := itemService.query.GetCategoryByID(categoryID)
	if categoryModel == nil {
		return nil
	}

	// transform data from model
	categoryDTO := itemService.transform.TransformGetCategoryByID(categoryModel)
	return categoryDTO
}

// CreateItemCategory ...
func (itemService *ItemService) CreateItemCategory(itemCategory *dto.ItemCategoryReq) (int64, error) {
	// transform dto to model
	itemCategoryCreate := itemService.transform.TransformCreateCategory(itemCategory)

	// add data to db
	id, err := itemService.command.CreateCategory(itemCategoryCreate)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateItemCategory ...
func (itemService *ItemService) UpdateItemCategory(categoryID int, itemCategory *dto.ItemCategoryReq) error {
	// transform dto to model
	itemCategoryUpdate := itemService.transform.TransformUpdateCategory(itemCategory)

	// update to db
	err := itemService.command.UpdateCategory(categoryID, itemCategoryUpdate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItemCategory ...
func (itemService *ItemService) DeleteItemCategory(categoryID int) error {
	// delete data from db
	err := itemService.command.DeleteCategory(categoryID)
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
