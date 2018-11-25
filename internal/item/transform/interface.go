package transform

import (
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
)

// ItemTransformInterface ...
type ItemTransformInterface interface {
	TransformGetCategory([]*model.ItemCategoryModelRead) []*dto.ItemCategoryIncRes
}
