package query

import "github.com/api-plastik/internal/item/model"

// ItemQueryInterface ...
type ItemQueryInterface interface {
	GetCategory() ([]*model.ItemCategoryModelRead, error)
	GetCategoryByID(categoryID int) (*model.ItemCategoryModelRead, error)
}
