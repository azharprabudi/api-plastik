package transform

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
)

// TransformCreateCategory ...
func (it *ItemTransform) TransformCreateCategory(category *dto.ItemCategoryReq) *model.ItemCategoryCreate {
	create := &model.ItemCategoryCreate{
		ItemCategory: model.ItemCategory{
			ItemCategoryID: uuid.NewV4(),
			Name:           category.Name,
			CreatedAt:      time.Now().UTC(),
		},
	}
	return create
}

// TransformUpdateCategory ...
func (it *ItemTransform) TransformUpdateCategory(category *dto.ItemCategoryReq) *model.ItemCategoryUpdate {
	update := &model.ItemCategoryUpdate{
		Name: category.Name,
	}
	return update
}

// TransformGetCategory ...
func (it *ItemTransform) TransformGetCategory(categories []*model.ItemCategoryRead) []*dto.ItemCategoryRes {
	// init variable
	var res = []*dto.ItemCategoryRes{}

	// transform data as dto expected
	for _, category := range categories {
		res = append(res, &dto.ItemCategoryRes{
			ItemCategoryID: category.ItemCategory.ItemCategoryID,
			ItemCategoryReq: dto.ItemCategoryReq{
				Name: category.ItemCategory.Name,
			},
			CreatedAt: category.ItemCategory.CreatedAt,
		})
	}

	return res
}

// TransformGetCategoryByID ...
func (it *ItemTransform) TransformGetCategoryByID(category *model.ItemCategoryRead) *dto.ItemCategoryRes {
	return &dto.ItemCategoryRes{
		ItemCategoryID: category.ItemCategoryID,
		ItemCategoryReq: dto.ItemCategoryReq{
			Name: category.ItemCategory.Name,
		},
		CreatedAt: category.ItemCategory.CreatedAt,
	}
}

// TransformCreateItem ...
func (it *ItemTransform) TransformCreateItem(item *dto.ItemReq) *model.ItemCreate {
	create := &model.ItemCreate{
		Item: model.Item{
			ItemID:         uuid.NewV4(),
			Name:           item.Name,
			ItemCategoryID: item.ItemCategoryID,
			CreatedAt:      time.Now().UTC(),
			UnitID:         item.UnitID,
		},
	}
	return create
}

// TransformUpdateItem ...
func (it *ItemTransform) TransformUpdateItem(item *dto.ItemReq) *model.ItemUpdate {
	update := &model.ItemUpdate{
		Name:           item.Name,
		ItemCategoryID: item.ItemCategoryID,
	}
	return update
}

// TransformGetItem ...
func (it *ItemTransform) TransformGetItem(items []*model.ItemRead) []*dto.ItemRes {
	// init variable
	var res = []*dto.ItemRes{}

	// transform data as dto expected
	for _, item := range items {
		res = append(res, &dto.ItemRes{
			ItemID:    item.Item.ItemID,
			CreatedAt: item.Item.CreatedAt,
			ItemReq: dto.ItemReq{
				Name:           item.Item.Name,
				ItemCategoryID: item.Item.ItemCategoryID,
				UnitID:         item.Item.UnitID,
			},
		})
	}

	return res
}

// TransformGetItemByID ...
func (it *ItemTransform) TransformGetItemByID(item *model.ItemRead) *dto.ItemRes {
	return &dto.ItemRes{
		ItemID:    item.Item.ItemID,
		CreatedAt: item.Item.CreatedAt,
		ItemReq: dto.ItemReq{
			Name:           item.Item.Name,
			ItemCategoryID: item.Item.ItemCategoryID,
			UnitID:         item.Item.UnitID,
		},
	}
}

// TransformGetItemUnit ...
func (it *ItemTransform) TransformGetItemUnit(itemUnit []*model.ItemUnitRead) []*dto.ItemUnitRes {
	var results []*dto.ItemUnitRes
	for _, unit := range itemUnit {
		results = append(results, &dto.ItemUnitRes{
			ItemUnit: dto.ItemUnit{
				ID:        unit.ID,
				Name:      unit.Name,
				CreatedAt: unit.CreatedAt,
			},
		})
	}
	return results
}

// NewItemTransform ...
func NewItemTransform() ItemTransformInterface {
	return ItemTransformSingleton
}
