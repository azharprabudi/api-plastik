package query

import "github.com/api-plastik/internal/item/model"

// ItemQueryInterface ...
type ItemQueryInterface interface {
	// Item Category
	GetCategory() ([]*model.ItemCategoryModelRead, error)
	GetCategoryByID(int) *model.ItemCategoryModelRead

	// Item
	GetItem() ([]*model.ItemRead, error)
	GetItemByID(string) *model.ItemRead
}
