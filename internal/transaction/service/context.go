package service

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/azharprabudi/api-plastik/internal/transaction/command"
	"github.com/azharprabudi/api-plastik/internal/transaction/event"
	"github.com/azharprabudi/api-plastik/internal/transaction/query"
	"github.com/azharprabudi/api-plastik/internal/transaction/transform"
)

// TransactionService ...
type TransactionService struct {
	db          *db.DB
	command     command.TransactionCommandInterface
	query       query.TransactionQueryInterface
	transform   transform.TransactionTransformInterface
	itemService service.ItemServiceInterface
	event       event.TransactionEventInterface
}
