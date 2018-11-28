package command

import "github.com/api-plastik/internal/item/model"

// ItemCommandInterface ...
type ItemCommandInterface interface {
	// item category
	CreateCategory(*model.ItemCategoryCreate) (int64, error)
	UpdateCategory(int, *model.ItemCategoryUpdate) error
	DeleteCategory(int) error

	// item
	CreateItem(*model.ItemCreate) error
	UpdateItem(string, *model.ItemUpdate) error
	DeleteItem(string) error
}
