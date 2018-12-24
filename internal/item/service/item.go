package service

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/azharprabudi/api-plastik/internal/item/transform"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/item/command"
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/query"
)

// GetItemCategories ...
func (is *ItemService) GetItemCategories(companyID uuid.UUID) ([]*dto.ItemCategoryRes, error) {
	itemCategories, err := is.query.GetCategories(companyID)
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetCategories(itemCategories), nil
}

// GetItemCategoryByID ...
func (is *ItemService) GetItemCategoryByID(companyID uuid.UUID, id uuid.UUID) (*dto.ItemCategoryRes, error) {
	itemCategory, err := is.query.GetCategoryByID(companyID, id)
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetCategoryByID(itemCategory), nil
}

// CreateItemCategory ...
func (is *ItemService) CreateItemCategory(companyID uuid.UUID, req *dto.ItemCategoryReq) (uuid.UUID, error) {
	itemCategory := is.transform.MakeModelCreateCategory(companyID, req)
	err := is.command.CreateCategory(itemCategory)
	if err != nil {
		return uuid.Nil, err
	}

	return itemCategory.ItemCategory.ItemCategoryID, nil
}

// UpdateItemCategory ...
func (is *ItemService) UpdateItemCategory(companyID uuid.UUID, id uuid.UUID, req *dto.ItemCategoryReq) error {
	itemCategory := is.transform.MakeModelUpdateCategory(req)
	err := is.command.UpdateCategory(companyID, id, itemCategory)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItemCategory ...
func (is *ItemService) DeleteItemCategory(companyID uuid.UUID, id uuid.UUID) error {
	err := is.command.DeleteCategory(companyID, id)
	if err != nil {
		return err
	}

	return nil
}

// GetItems ...
func (is *ItemService) GetItems(companyID uuid.UUID) ([]*dto.ItemRes, error) {
	items, err := is.query.GetItems(companyID)
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetItems(items), nil
}

// GetItemByID ...
func (is *ItemService) GetItemByID(companyID uuid.UUID, id uuid.UUID) (*dto.ItemRes, error) {
	item, err := is.query.GetItemByID(companyID, id)
	if err != nil {
		return nil, err
	}

	return is.transform.MakeResponseGetItemByID(item), nil
}

// CreateItem ...
func (is *ItemService) CreateItem(companyID uuid.UUID, req *dto.ItemReq) (uuid.UUID, error) {
	item := is.transform.MakeModelCreateItem(companyID, req)
	err := is.command.CreateItem(item)
	if err != nil {
		return uuid.Nil, err
	}

	return item.Item.ItemID, nil
}

// UpdateItem ...
func (is *ItemService) UpdateItem(companyID uuid.UUID, id uuid.UUID, req *dto.ItemReq) error {
	item := is.transform.MakeModelUpdateItem(req)
	err := is.command.UpdateItem(companyID, id, item)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem ...
func (is *ItemService) DeleteItem(companyID uuid.UUID, id uuid.UUID) error {
	err := is.command.DeleteItem(companyID, id)
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

// CreateItemStockLog ...
func (is *ItemService) CreateItemStockLog(tx *sqlx.Tx, item *model.ItemStockLogCreate) error {
	err := is.command.CreateItemStockLog(tx, item)
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
