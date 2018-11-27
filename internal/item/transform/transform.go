package transform

import (
	"time"

	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
)

// TransformCreateCategory ...
func (it *ItemTransform) TransformCreateCategory(itemCategoryModelDTO *dto.ItemCategoryIncReq) *model.ItemCategoryCreate {
	itemCategoryCreate := &model.ItemCategoryCreate{
		Name:      itemCategoryModelDTO.Name,
		CreatedAt: time.Now().UTC(),
	}
	return itemCategoryCreate
}

// TransformUpdateCategory ...
func (it *ItemTransform) TransformUpdateCategory(itemCategoryModelDTO *dto.ItemCategoryIncReq) *model.ItemCategoryUpdate {
	itemCategoryUpdate := &model.ItemCategoryUpdate{
		Name: itemCategoryModelDTO.Name,
	}
	return itemCategoryUpdate
}

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

// TransformGetCategoryByID ...
func (it *ItemTransform) TransformGetCategoryByID(itemCategoryModelRead *model.ItemCategoryModelRead) *dto.ItemCategoryIncRes {
	return &dto.ItemCategoryIncRes{
		ID:        itemCategoryModelRead.ItemCategoryID.ItemCategoryID,
		Name:      itemCategoryModelRead.Name,
		CreatedAt: itemCategoryModelRead.CreatedAt,
	}
}

// NewItemTransform ...
func NewItemTransform() ItemTransformInterface {
	return ItemTransformSingleton
}
