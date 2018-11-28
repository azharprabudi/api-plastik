package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/model"
)

// TransformCreateCategory ...
func (it *ItemTransform) TransformCreateCategory(itemCategoryModelDTO *dto.ItemCategoryReq) *model.ItemCategoryCreate {
	itemCategoryCreate := &model.ItemCategoryCreate{
		Name:      itemCategoryModelDTO.Name,
		CreatedAt: time.Now().UTC(),
	}
	return itemCategoryCreate
}

// TransformUpdateCategory ...
func (it *ItemTransform) TransformUpdateCategory(itemCategoryModelDTO *dto.ItemCategoryReq) *model.ItemCategoryUpdate {
	itemCategoryUpdate := &model.ItemCategoryUpdate{
		Name: itemCategoryModelDTO.Name,
	}
	return itemCategoryUpdate
}

// TransformGetCategory ...
func (it *ItemTransform) TransformGetCategory(itemCategoryModelRead []*model.ItemCategoryModelRead) []*dto.ItemCategoryRes {
	// init variable
	var itemCategoryIncRes = []*dto.ItemCategoryRes{}

	// transform data as dto expected
	for _, item := range itemCategoryModelRead {
		itemCategoryIncRes = append(itemCategoryIncRes, &dto.ItemCategoryRes{
			ID:        item.ItemCategoryID.ItemCategoryID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
		})
	}

	return itemCategoryIncRes
}

// TransformGetCategoryByID ...
func (it *ItemTransform) TransformGetCategoryByID(itemCategoryModelRead *model.ItemCategoryModelRead) *dto.ItemCategoryRes {
	return &dto.ItemCategoryRes{
		ID:        itemCategoryModelRead.ItemCategoryID.ItemCategoryID,
		Name:      itemCategoryModelRead.Name,
		CreatedAt: itemCategoryModelRead.CreatedAt,
	}
}

// TransformCreateItem ...
func (it *ItemTransform) TransformCreateItem(itemDTO *dto.ItemReq) *model.ItemCreate {
	itemCreate := &model.ItemCreate{
		ItemID:         uuid.NewV4(),
		Name:           itemDTO.Name,
		CategoryItemID: itemDTO.CategoryID,
		CreatedAt:      time.Now().UTC(),
	}
	return itemCreate
}

// TransformUpdateItem ...
func (it *ItemTransform) TransformUpdateItem(itemDTO *dto.ItemReq) *model.ItemUpdate {
	itemUpdate := &model.ItemUpdate{
		Name:           itemDTO.Name,
		CategoryItemID: itemDTO.CategoryID,
	}
	return itemUpdate
}

// TransformGetItem ...
func (it *ItemTransform) TransformGetItem(itemRead []*model.ItemRead) []*dto.ItemRes {
	// init variable
	var itemRes = []*dto.ItemRes{}

	// transform data as dto expected
	for _, item := range itemRead {
		itemRes = append(itemRes, &dto.ItemRes{
			ID:        item.ItemCreate.ItemID,
			CreatedAt: item.ItemCreate.CreatedAt,
			ItemReq: dto.ItemReq{
				Name:       item.ItemCreate.Name,
				CategoryID: item.ItemCreate.CategoryItemID,
			},
		})
	}

	return itemRes
}

// TransformGetItemByID ...
func (it *ItemTransform) TransformGetItemByID(itemRead *model.ItemRead) *dto.ItemRes {
	return &dto.ItemRes{
		ID:        itemRead.ItemCreate.ItemID,
		CreatedAt: itemRead.ItemCreate.CreatedAt,
		ItemReq: dto.ItemReq{
			Name:       itemRead.ItemCreate.Name,
			CategoryID: itemRead.ItemCreate.CategoryItemID,
		},
	}
}

// NewItemTransform ...
func NewItemTransform() ItemTransformInterface {
	return ItemTransformSingleton
}
