package command

import (
	"github.com/api-plastik/internal/item/model"
	uuid "github.com/satori/go.uuid"
)

// ItemCommandInterface ...
type ItemCommandInterface interface {
	// item category
	CreateCategory(*model.ItemCategoryCreate) error
	UpdateCategory(uuid.UUID, *model.ItemCategoryUpdate) error
	DeleteCategory(uuid.UUID) error

	// item
	CreateItem(*model.ItemCreate) error
	UpdateItem(uuid.UUID, *model.ItemUpdate) error
	DeleteItem(uuid.UUID) error
}
