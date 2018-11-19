package helpers

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type key string

// KeyDB ...
const (
	KeyDB key = "DB"
)

// StartTransaction ...
func StartTransaction(db *sqlx.DB) context.Context {
	c := context.Background()
	tx, _ := db.Beginx()
	c = context.WithValue(c, KeyDB, tx)
	return c
}

// RunTransaction ...
func RunTransaction(context context.Context, function func(*sqlx.Tx) error) error {
	tx := context.Value(KeyDB).(sqlx.Tx)
	err := function(&tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
