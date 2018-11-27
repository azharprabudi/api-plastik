package service

import "github.com/api-plastik/internal/item/dto"

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(*dto.ItemCategoryReq) error
	UpdateItemCategory(int, *dto.ItemCategoryReq) error
	DeleteItemCategory(int) error
	GetItemCategory() ([]*dto.ItemCategoryRes, error)
	GetItemCategoryByID(int) *dto.ItemCategoryRes
}
