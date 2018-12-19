package query

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	uuid "github.com/satori/go.uuid"
)

// ItemQueryInterface ...
type ItemQueryInterface interface {
	// Item Category
	GetCategory() ([]*model.ItemCategoryRead, error)
	GetCategoryByID(uuid.UUID) *model.ItemCategoryRead

	// Item
	GetItem() ([]*model.ItemRead, error)
	GetItemByID(uuid.UUID) *model.ItemRead

	// Item Unit
	GetItemUnit() ([]*model.ItemUnitRead, error)
}
