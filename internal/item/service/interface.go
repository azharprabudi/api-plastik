package service

import (
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	uuid "github.com/satori/go.uuid"
)

// ItemServiceInterface ...
type ItemServiceInterface interface {
	// Category
	CreateItemCategory(*dto.ItemCategoryReq) (uuid.UUID, error)
	UpdateItemCategory(uuid.UUID, *dto.ItemCategoryReq) error
	DeleteItemCategory(uuid.UUID) error
	GetItemCategory() ([]*dto.ItemCategoryRes, error)
	GetItemCategoryByID(uuid.UUID) *dto.ItemCategoryRes

	// Item
	CreateItem(*dto.ItemReq) (uuid.UUID, error)
	UpdateItem(uuid.UUID, *dto.ItemReq) error
	DeleteItem(uuid.UUID) error
	GetItem() ([]*dto.ItemRes, error)
	GetItemByID(uuid.UUID) *dto.ItemRes

	// Item Unit
	GetItemUnit() ([]*dto.ItemUnitRes, error)
}
