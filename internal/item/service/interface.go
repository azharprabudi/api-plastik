package service

import "github.com/api-plastik/internal/item/dto"

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(*dto.ItemCategoryIncReq) error
	UpdateItemCategory(int, *dto.ItemCategoryIncReq) error
	DeleteItemCategory(int) error
	GetItemCategory() ([]*dto.ItemCategoryIncRes, error)
	GetItemCategoryByID(int) (*dto.ItemCategoryIncRes, error)
}
