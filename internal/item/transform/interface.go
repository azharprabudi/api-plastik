package transform

import (
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/model"
	uuid "github.com/satori/go.uuid"
)

// ItemTransformInterface ...
type ItemTransformInterface interface {
	MakeModelCreateCategory(uuid.UUID, *dto.ItemCategoryReq) *model.ItemCategoryCreate
	MakeModelUpdateCategory(*dto.ItemCategoryReq) *model.ItemCategoryUpdate
	MakeResponseGetCategories([]*model.ItemCategoryRead) []*dto.ItemCategoryRes
	MakeResponseGetCategoryByID(*model.ItemCategoryRead) *dto.ItemCategoryRes
	MakeModelCreateItem(uuid.UUID, *dto.ItemReq) *model.ItemCreate
	MakeModelUpdateItem(*dto.ItemReq) *model.ItemUpdate
	MakeResponseGetItems([]*model.ItemRead) []*dto.ItemRes
	MakeResponseGetItemByID(*model.ItemRead) *dto.ItemRes
	MakeResponseGetItemUnits([]*model.ItemUnitRead) []*dto.ItemUnitRes
	MakeResponseGetItemStockLogs([]*model.ItemStockLogRead) []*dto.ItemStockLogRes
}
