package trx

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// TransactionInterface ...
type TransactionInterface interface {
	CreateTrx(db *sqlx.DB) context.Context
	RunTrx(ctx context.Context, cb func(tx *sqlx.Tx) error) error
}
