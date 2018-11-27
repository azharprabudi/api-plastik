package transform

import (
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
)

// ItemTransformInterface ...
type ItemTransformInterface interface {
	TransformCreateCategory(*dto.ItemCategoryIncReq) *model.ItemCategoryCreate
	TransformUpdateCategory(*dto.ItemCategoryIncReq) *model.ItemCategoryUpdate
	TransformGetCategory([]*model.ItemCategoryModelRead) []*dto.ItemCategoryIncRes
	TransformGetCategoryByID(*model.ItemCategoryModelRead) *dto.ItemCategoryIncRes
}
