package service

import (
	"github.com/azharprabudi/api-plastik/internal/item/transform"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/item/command"
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/query"
)

// GetItemCategory ...
func (is *ItemService) GetItemCategory() ([]*dto.ItemCategoryRes, error) {
	// add data to db
	categories, err := is.query.GetCategory()
	if err != nil {
		return nil, err
	}

	// transform data from model
	categoriesDTO := is.transform.TransformGetCategory(categories)
	return categoriesDTO, nil
}

// GetItemCategoryByID ...
func (is *ItemService) GetItemCategoryByID(itemCatID uuid.UUID) *dto.ItemCategoryRes {
	category := is.query.GetCategoryByID(itemCatID)
	if category == nil {
		return nil
	}

	// transform data from model
	categoryDTO := is.transform.TransformGetCategoryByID(category)
	return categoryDTO
}

// CreateItemCategory ...
func (is *ItemService) CreateItemCategory(itemCategory *dto.ItemCategoryReq) (uuid.UUID, error) {
	// transform dto to model
	create := is.transform.TransformCreateCategory(itemCategory)

	// add data to db
	err := is.command.CreateCategory(create)
	if err != nil {
		return uuid.Nil, err
	}

	return create.ItemCategoryID, nil
}

// UpdateItemCategory ...
func (is *ItemService) UpdateItemCategory(itemCatID uuid.UUID, itemCategory *dto.ItemCategoryReq) error {
	// transform dto to model
	update := is.transform.TransformUpdateCategory(itemCategory)

	// update to db
	err := is.command.UpdateCategory(itemCatID, update)
	if err != nil {
		return err
	}
	return nil
}

// DeleteItemCategory ...
func (is *ItemService) DeleteItemCategory(itemCatID uuid.UUID) error {
	// delete data from db
	err := is.command.DeleteCategory(itemCatID)
	if err != nil {
		return err
	}
	return nil
}

// GetItem ...
func (is *ItemService) GetItem() ([]*dto.ItemRes, error) {
	// add data to db
	item, err := is.query.GetItem()
	if err != nil {
		return nil, err
	}

	// transform data from model
	itemDTO := is.transform.TransformGetItem(item)
	return itemDTO, nil
}

// GetItemByID ...
func (is *ItemService) GetItemByID(itemID uuid.UUID) *dto.ItemRes {
	item := is.query.GetItemByID(itemID)
	if item == nil {
		return nil
	}

	// transform data from model
	itemDTO := is.transform.TransformGetItemByID(item)
	return itemDTO
}

// CreateItem ...
func (is *ItemService) CreateItem(item *dto.ItemReq) (uuid.UUID, error) {
	// transform dto to model
	create := is.transform.TransformCreateItem(item)

	// add data to db
	err := is.command.CreateItem(create)
	if err != nil {
		return uuid.Nil, err
	}

	return create.Item.ItemID, nil
}

// UpdateItem ...
func (is *ItemService) UpdateItem(itemID uuid.UUID, item *dto.ItemReq) error {
	// transform dto to model
	update := is.transform.TransformUpdateItem(item)

	// update to db
	err := is.command.UpdateItem(itemID, update)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem ...
func (is *ItemService) DeleteItem(itemID uuid.UUID) error {
	// delete data from db
	err := is.command.DeleteItem(itemID)
	if err != nil {
		return err
	}

	return nil
}

// GetItemUnit ...
func (is *ItemService) GetItemUnit() ([]*dto.ItemUnitRes, error) {
	// get from tbl item unit
	res, err := is.query.GetItemUnit()
	if err != nil {
		return nil, err
	}

	itemUnitDTO := is.transform.TransformGetItemUnit(res)
	return itemUnitDTO, nil
}

// NewItemService ...
func NewItemService(db *db.DB) ItemServiceInterface {
	return &ItemService{
		query:     query.NewItemQuery(db),
		command:   command.NewItemCommand(db),
		transform: transform.NewItemTransform(),
	}
}
