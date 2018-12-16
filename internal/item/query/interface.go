package query

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/satori/go.uuid"
)

// ItemQueryInterface ...
type ItemQueryInterface interface {
	// Item Category
	GetCategory() ([]*model.ItemCategoryRead, error)
	GetCategoryByID(uuid.UUID) *model.ItemCategoryRead

	// Item
	GetItem() ([]*model.ItemRead, error)
	GetItemByID(uuid.UUID) *model.ItemRead
}
