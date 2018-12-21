package transform

import (
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
)

// ItemTransformInterface ...
type ItemTransformInterface interface {
	// category
	MakeModelCreateCategory(*dto.ItemCategoryReq) *model.ItemCategoryCreate
	MakeModelUpdateCategory(*dto.ItemCategoryReq) *model.ItemCategoryUpdate
	MakeResponseGetCategories([]*model.ItemCategoryRead) []*dto.ItemCategoryRes
	MakeResponseGetCategoryByID(*model.ItemCategoryRead) *dto.ItemCategoryRes

	// item
	MakeModelCreateItem(*dto.ItemReq) *model.ItemCreate
	MakeModelUpdateItem(*dto.ItemReq) *model.ItemUpdate
	MakeResponseGetItems([]*model.ItemRead) []*dto.ItemRes
	MakeResponseGetItemByID(*model.ItemRead) *dto.ItemRes

	// item unit
	MakeResponseGetItemUnits([]*model.ItemUnitRead) []*dto.ItemUnitRes
}
