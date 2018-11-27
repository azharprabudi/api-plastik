package command

import "github.com/api-plastik/internal/item/model"

// ItemCommandInterface ...
type ItemCommandInterface interface {
	CreateCategory(*model.ItemCategoryCreate) (int64, error)
	UpdateCategory(int, *model.ItemCategoryUpdate) error
	DeleteCategory(int) error
}
