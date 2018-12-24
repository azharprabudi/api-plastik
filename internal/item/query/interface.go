package query

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	uuid "github.com/satori/go.uuid"
)

// ItemQueryInterface ...
type ItemQueryInterface interface {
	GetCategories(uuid.UUID) ([]*model.ItemCategoryRead, error)
	GetCategoryByID(uuid.UUID, uuid.UUID) (*model.ItemCategoryRead, error)
	GetItems(uuid.UUID) ([]*model.ItemRead, error)
	GetItemByID(uuid.UUID, uuid.UUID) (*model.ItemRead, error)
	GetItemUnits() ([]*model.ItemUnitRead, error)
}
