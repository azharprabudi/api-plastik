package command

import "github.com/api-plastik/internal/item/model"

// ItemCommandInterface ...
type ItemCommandInterface interface {
	CreateCategory(*model.ItemCategoryCreate) error
}
