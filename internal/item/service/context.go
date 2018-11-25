package service

import (
	"github.com/api-plastik/internal/item/command"
	"github.com/api-plastik/internal/item/query"
	"github.com/api-plastik/internal/item/transform"
)

// ItemService ...
type ItemService struct {
	query     query.ItemQueryInterface
	command   command.ItemCommandInterface
	transform transform.ItemTransformInterface
}
