package service

import (
	"github.com/api-plastik/internal/item/dto"
	"github.com/satori/go.uuid"
)

// ItemServiceInterface ...
type ItemServiceInterface interface {
	// Category
	CreateItemCategory(*dto.ItemCategoryReq) (int64, error)
	UpdateItemCategory(int, *dto.ItemCategoryReq) error
	DeleteItemCategory(int) error
	GetItemCategory() ([]*dto.ItemCategoryRes, error)
	GetItemCategoryByID(int) *dto.ItemCategoryRes

	// Item
	CreateItem(*dto.ItemReq) (uuid.UUID, error)
	UpdateItem(string, *dto.ItemReq) error
	DeleteItem(string) error
	GetItem() ([]*dto.ItemRes, error)
	GetItemByID(string) *dto.ItemRes
}
