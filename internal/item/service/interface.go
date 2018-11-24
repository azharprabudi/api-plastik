package service

import "github.com/api-plastik/internal/item/dto"

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(*dto.ItemCategoryIncReq) error
}
