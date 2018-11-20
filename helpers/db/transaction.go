package helpers

import (
	"context"

	"github.com/api-plastik/config"
	"github.com/jmoiron/sqlx"
)

// StartTransaction ...
func StartTransaction(db *sqlx.DB) context.Context {
	c := context.Background()
	tx, _ := db.Beginx()
	c = context.WithValue(c, config.DBKey, tx)
	return c
}

// RunTransaction ...
func RunTransaction(context context.Context, function func(*sqlx.Tx) error) error {
	tx := context.Value(config.DBKey).(*sqlx.Tx)
	err := function(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
