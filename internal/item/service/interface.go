package service

import (
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(*dto.ItemCategoryReq) (uuid.UUID, error)
	UpdateItemCategory(uuid.UUID, *dto.ItemCategoryReq) error
	DeleteItemCategory(uuid.UUID) error
	GetItemCategories() ([]*dto.ItemCategoryRes, error)
	GetItemCategoryByID(uuid.UUID) (*dto.ItemCategoryRes, error)
	CreateItem(*dto.ItemReq) (uuid.UUID, error)
	UpdateItem(uuid.UUID, *dto.ItemReq) error
	DeleteItem(uuid.UUID) error
	GetItems() ([]*dto.ItemRes, error)
	GetItemByID(uuid.UUID) (*dto.ItemRes, error)
	GetItemUnits() ([]*dto.ItemUnitRes, error)
	CreateItemStockLog(*sqlx.Tx, *model.ItemStockLogCreate) error
}
