package service

import (
	"github.com/api-plastik/internal/item/transform"
	"github.com/satori/go.uuid"

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

// GetItem ...
func (itemService *ItemService) GetItem() ([]*dto.ItemRes, error) {
	// add data to db
	itemModel, err := itemService.query.GetItem()
	if err != nil {
		return nil, err
	}

	// transform data from model
	itemDTO := itemService.transform.TransformGetItem(itemModel)
	return itemDTO, nil
}

// GetItemByID ...
func (itemService *ItemService) GetItemByID(itemID string) *dto.ItemRes {
	itemModel := itemService.query.GetItemByID(itemID)
	if itemModel == nil {
		return nil
	}

	// transform data from model
	itemDTO := itemService.transform.TransformGetItemByID(itemModel)
	return itemDTO
}

// CreateItem ...
func (itemService *ItemService) CreateItem(item *dto.ItemReq) (uuid.UUID, error) {
	// transform dto to model
	itemCreate := itemService.transform.TransformCreateItem(item)

	// add data to db
	err := itemService.command.CreateItem(itemCreate)
	if err != nil {
		return uuid.Nil, err
	}
	return itemCreate.ItemID, nil
}

// UpdateItem ...
func (itemService *ItemService) UpdateItem(itemID string, item *dto.ItemReq) error {
	// transform dto to model
	itemUpdate := itemService.transform.TransformUpdateItem(item)

	// update to db
	err := itemService.command.UpdateItem(itemID, itemUpdate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItem ...
func (itemService *ItemService) DeleteItem(itemID string) error {
	// delete data from db
	err := itemService.command.DeleteItem(itemID)
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
