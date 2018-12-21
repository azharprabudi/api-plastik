package command

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	uuid "github.com/satori/go.uuid"
)

// ItemCommandInterface ...
type ItemCommandInterface interface {
	CreateCategory(*model.ItemCategoryCreate) error
	UpdateCategory(uuid.UUID, *model.ItemCategoryUpdate) error
	DeleteCategory(uuid.UUID) error
	CreateItem(*model.ItemCreate) error
	UpdateItem(uuid.UUID, *model.ItemUpdate) error
	DeleteItem(uuid.UUID) error
}
