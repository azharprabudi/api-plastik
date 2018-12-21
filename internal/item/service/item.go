package service

import (
	"github.com/azharprabudi/api-plastik/internal/item/transform"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/item/command"
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/query"
)

// GetItemCategories ...
func (is *ItemService) GetItemCategories() ([]*dto.ItemCategoryRes, error) {
	itemCategories, err := is.query.GetCategories()
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetCategories(itemCategories), nil
}

// GetItemCategoryByID ...
func (is *ItemService) GetItemCategoryByID(id uuid.UUID) (*dto.ItemCategoryRes, error) {
	itemCategory, err := is.query.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetCategoryByID(itemCategory), nil
}

// CreateItemCategory ...
func (is *ItemService) CreateItemCategory(req *dto.ItemCategoryReq) (uuid.UUID, error) {
	itemCategory := is.transform.MakeModelCreateCategory(req)
	err := is.command.CreateCategory(itemCategory)
	if err != nil {
		return uuid.Nil, err
	}

	return itemCategory.ItemCategory.ItemCategoryID, nil
}

// UpdateItemCategory ...
func (is *ItemService) UpdateItemCategory(id uuid.UUID, req *dto.ItemCategoryReq) error {
	itemCategory := is.transform.MakeModelUpdateCategory(req)
	err := is.command.UpdateCategory(id, itemCategory)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItemCategory ...
func (is *ItemService) DeleteItemCategory(id uuid.UUID) error {
	err := is.command.DeleteCategory(id)
	if err != nil {
		return err
	}

	return nil
}

// GetItems ...
func (is *ItemService) GetItems() ([]*dto.ItemRes, error) {
	items, err := is.query.GetItems()
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetItems(items), nil
}

// GetItemByID ...
func (is *ItemService) GetItemByID(id uuid.UUID) (*dto.ItemRes, error) {
	item, err := is.query.GetItemByID(id)
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetItemByID(item), nil
}

// CreateItem ...
func (is *ItemService) CreateItem(req *dto.ItemReq) (uuid.UUID, error) {
	item := is.transform.MakeModelCreateItem(req)
	err := is.command.CreateItem(item)
	if err != nil {
		return uuid.Nil, err
	}

	return item.Item.ItemID, nil
}

// UpdateItem ...
func (is *ItemService) UpdateItem(id uuid.UUID, req *dto.ItemReq) error {
	item := is.transform.MakeModelUpdateItem(req)
	err := is.command.UpdateItem(id, item)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem ...
func (is *ItemService) DeleteItem(id uuid.UUID) error {
	err := is.command.DeleteItem(id)
	if err != nil {
		return err
	}

	return nil
}

// GetItemUnits ...
func (is *ItemService) GetItemUnits() ([]*dto.ItemUnitRes, error) {
	res, err := is.query.GetItemUnits()
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetItemUnits(res), nil
}

// NewItemService ...
func NewItemService(db *db.DB) ItemServiceInterface {
	return &ItemService{
		query:     query.NewItemQuery(db),
		command:   command.NewItemCommand(db),
		transform: transform.NewItemTransform(),
	}
}
