package event

import (
	"github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/jmoiron/sqlx"
)

// TransactionEventInterface ...
type TransactionEventInterface interface {
	TriggerAfterCreateTransaction(*sqlx.Tx, []*model.ItemStockLogCreate) error
}
