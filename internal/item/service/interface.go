package service

import "github.com/api-plastik/dto"

// ItemServiceInterface ...
type ItemServiceInterface interface {
	CreateItemCategory(*dto.ItemCategoryIncReq) error
}
