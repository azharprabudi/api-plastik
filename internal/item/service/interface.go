package service

import (
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(uuid.UUID, *dto.ItemCategoryReq) (uuid.UUID, error)
	UpdateItemCategory(uuid.UUID, uuid.UUID, *dto.ItemCategoryReq) error
	DeleteItemCategory(uuid.UUID, uuid.UUID) error
	GetItemCategories(uuid.UUID) ([]*dto.ItemCategoryRes, error)
	GetItemCategoryByID(uuid.UUID, uuid.UUID) (*dto.ItemCategoryRes, error)
	CreateItem(uuid.UUID, *dto.ItemReq) (uuid.UUID, error)
	UpdateItem(uuid.UUID, uuid.UUID, *dto.ItemReq) error
	DeleteItem(uuid.UUID, uuid.UUID) error
	GetItems(uuid.UUID) ([]*dto.ItemRes, error)
	GetItemByID(uuid.UUID, uuid.UUID) (*dto.ItemRes, error)
	GetItemUnits() ([]*dto.ItemUnitRes, error)
	CreateItemStockLog(*sqlx.Tx, *model.ItemStockLogCreate) error
	GetItemStockLogs(uuid.UUID) ([]*dto.ItemStockLogRes, error)
}
