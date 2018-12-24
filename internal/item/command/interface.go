package command

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// ItemCommandInterface ...
type ItemCommandInterface interface {
	CreateCategory(*model.ItemCategoryCreate) error
	UpdateCategory(uuid.UUID, uuid.UUID, *model.ItemCategoryUpdate) error
	DeleteCategory(uuid.UUID, uuid.UUID) error
	CreateItem(*model.ItemCreate) error
	UpdateItem(uuid.UUID, uuid.UUID, *model.ItemUpdate) error
	DeleteItem(uuid.UUID, uuid.UUID) error
	CreateItemStockLog(*sqlx.Tx, *model.ItemStockLogCreate) error
}
