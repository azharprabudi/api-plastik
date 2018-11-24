package service

import (
	"github.com/api-plastik/internal/item/command"
	"github.com/api-plastik/internal/item/query"
)

// ItemService ...
type ItemService struct {
	query   query.ItemQueryInterface
	command command.ItemCommandInterface
}
