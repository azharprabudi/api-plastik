package trx

import (
	"context"
	"errors"

	"github.com/azharprabudi/api-plastik/config"
	"github.com/jmoiron/sqlx"
)

// CreateTrx ...
func (trx *Transaction) CreateTrx(db *sqlx.DB) context.Context {
	c := context.Background()
	tx, _ := db.Beginx()
	c = context.WithValue(c, config.DBKey, tx)
	return c
}

// RunTrx ...
func (trx *Transaction) RunTrx(ctx context.Context, cb func(tx *sqlx.Tx) error) error {
	rawTx := ctx.Value(config.DBKey)
	if rawTx != nil {
		tx := rawTx.(*sqlx.Tx)
		err := cb(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}
	return errors.New("Transaction db not found")
}

// NewTransaction ...
func NewTransaction() TransactionInterface {
	return TransactionSingleton
}
