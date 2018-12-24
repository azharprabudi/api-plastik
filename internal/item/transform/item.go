package transform

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
)

// MakeModelCreateCategory ...
func (it *ItemTransform) MakeModelCreateCategory(companyID uuid.UUID, req *dto.ItemCategoryReq) *model.ItemCategoryCreate {
	return &model.ItemCategoryCreate{
		ItemCategory: model.ItemCategory{
			ItemCategoryID: uuid.NewV4(),
			Name:           req.Name,
			CreatedAt:      time.Now().UTC(),
			CompanyID:      companyID,
		},
	}
}

// MakeModelUpdateCategory ...
func (it *ItemTransform) MakeModelUpdateCategory(req *dto.ItemCategoryReq) *model.ItemCategoryUpdate {
	return &model.ItemCategoryUpdate{
		Name: req.Name,
	}
}

// MakeResponseGetCategories ...
func (it *ItemTransform) MakeResponseGetCategories(res []*model.ItemCategoryRead) []*dto.ItemCategoryRes {
	var results []*dto.ItemCategoryRes
	for _, category := range res {
		results = append(results, &dto.ItemCategoryRes{
			ItemCategoryID: category.ItemCategory.ItemCategoryID,
			ItemCategoryReq: dto.ItemCategoryReq{
				Name: category.ItemCategory.Name,
			},
			CreatedAt: category.ItemCategory.CreatedAt,
		})
	}

	return results
}

// MakeResponseGetCategoryByID ...
func (it *ItemTransform) MakeResponseGetCategoryByID(res *model.ItemCategoryRead) *dto.ItemCategoryRes {
	return &dto.ItemCategoryRes{
		ItemCategoryID: res.ItemCategoryID,
		ItemCategoryReq: dto.ItemCategoryReq{
			Name: res.ItemCategory.Name,
		},
		CreatedAt: res.ItemCategory.CreatedAt,
	}
}

// MakeModelCreateItem ...
func (it *ItemTransform) MakeModelCreateItem(companyID uuid.UUID, req *dto.ItemReq) *model.ItemCreate {
	return &model.ItemCreate{
		Item: model.Item{
			ItemID:         uuid.NewV4(),
			Name:           req.Name,
			ItemCategoryID: req.ItemCategoryID,
			CreatedAt:      time.Now().UTC(),
			UnitID:         req.UnitID,
			CompanyID:      companyID,
		},
	}
}

// MakeModelUpdateItem ...
func (it *ItemTransform) MakeModelUpdateItem(item *dto.ItemReq) *model.ItemUpdate {
	return &model.ItemUpdate{
		Name:           item.Name,
		ItemCategoryID: item.ItemCategoryID,
	}
}

// MakeResponseGetItems ...
func (it *ItemTransform) MakeResponseGetItems(res []*model.ItemRead) []*dto.ItemRes {
	var results = []*dto.ItemRes{}
	for _, item := range res {
		results = append(results, &dto.ItemRes{
			ItemID:    item.Item.ItemID,
			CreatedAt: item.Item.CreatedAt,
			Item: dto.Item{
				Name:           item.Item.Name,
				ItemCategoryID: item.Item.ItemCategoryID,
				UnitID:         item.Item.UnitID,
			},
		})
	}

	return results
}

// MakeResponseGetItemByID ...
func (it *ItemTransform) MakeResponseGetItemByID(res *model.ItemRead) *dto.ItemRes {
	return &dto.ItemRes{
		ItemID:    res.Item.ItemID,
		CreatedAt: res.Item.CreatedAt,
		Item: dto.Item{
			Name:           res.Item.Name,
			ItemCategoryID: res.Item.ItemCategoryID,
			UnitID:         res.Item.UnitID,
		},
	}
}

// MakeResponseGetItemUnits ...
func (it *ItemTransform) MakeResponseGetItemUnits(res []*model.ItemUnitRead) []*dto.ItemUnitRes {
	var results []*dto.ItemUnitRes
	for _, unit := range res {
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
