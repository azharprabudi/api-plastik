package transform

import (
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
)

// ItemTransformInterface ...
type ItemTransformInterface interface {
	// category
	TransformCreateCategory(*dto.ItemCategoryReq) *model.ItemCategoryCreate
	TransformUpdateCategory(*dto.ItemCategoryReq) *model.ItemCategoryUpdate
	TransformGetCategory([]*model.ItemCategoryRead) []*dto.ItemCategoryRes
	TransformGetCategoryByID(*model.ItemCategoryRead) *dto.ItemCategoryRes

	// item
	TransformCreateItem(*dto.ItemReq) *model.ItemCreate
	TransformUpdateItem(*dto.ItemReq) *model.ItemUpdate
	TransformGetItem([]*model.ItemRead) []*dto.ItemRes
	TransformGetItemByID(*model.ItemRead) *dto.ItemRes
}
