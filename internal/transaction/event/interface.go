package event

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// TransactionEventInterface ...
type TransactionEventInterface interface {
	TriggerAfterCreateTransaction(*sqlx.Tx, []*model.ItemStockLogCreate, uuid.UUID) error
}
