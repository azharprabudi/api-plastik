package service

import (
	"github.com/api-plastik/internal/item/dto"
	"github.com/satori/go.uuid"
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
}
