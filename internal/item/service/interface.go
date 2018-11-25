package service

import "github.com/api-plastik/internal/item/dto"

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(*dto.ItemCategoryIncReq) error
	GetItemCategory() ([]*dto.ItemCategoryIncRes, error)
	GetItemCategoryByID(categoryID int) (*dto.ItemCategoryIncRes, error)
}
