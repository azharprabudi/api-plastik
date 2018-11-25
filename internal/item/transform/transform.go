package transform

import (
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
)

// TransformGetCategory ...
func (it *ItemTransform) TransformGetCategory(itemCategoryModelRead []*model.ItemCategoryModelRead) []*dto.ItemCategoryIncRes {
	// init variable
	var itemCategoryIncRes = []*dto.ItemCategoryIncRes{}

	// transform data as dto expected
	for _, item := range itemCategoryModelRead {
		itemCategoryIncRes = append(itemCategoryIncRes, &dto.ItemCategoryIncRes{
			ID:        item.ItemCategoryID.ItemCategoryID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
		})
	}

	return itemCategoryIncRes
}

// NewItemTransform ...
func NewItemTransform() ItemTransformInterface {
	return ItemTransformSingleton
}
