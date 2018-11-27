package transform

import (
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
)

// ItemTransformInterface ...
type ItemTransformInterface interface {
	TransformCreateCategory(*dto.ItemCategoryReq) *model.ItemCategoryCreate
	TransformUpdateCategory(*dto.ItemCategoryReq) *model.ItemCategoryUpdate
	TransformGetCategory([]*model.ItemCategoryModelRead) []*dto.ItemCategoryRes
	TransformGetCategoryByID(*model.ItemCategoryModelRead) *dto.ItemCategoryRes
}
