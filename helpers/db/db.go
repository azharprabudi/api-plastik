package helpers

import (
	"context"
	"errors"
	"fmt"

	"github.com/api-plastik/config"
	"github.com/api-plastik/helpers/db/model"
	"github.com/jmoiron/sqlx"
)

/*
*
* this query builder, just working on sql db. Not compatible for
* nosql db
 */

// CreateTrx ...
func CreateTrx(db *sqlx.DB) context.Context {
	c := context.Background()
	tx, _ := db.Beginx()
	c = context.WithValue(c, config.DBKey, tx)
	return c
}

// RunTrx ...
func RunTrx(ctx context.Context, cb func(tx *sqlx.Tx) error) error {
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

// Query ...
func Query(tableName string, limit int, offset int) string {
	query := fmt.Sprintf("select * from %s", tableName)
	if limit != 0 && offset == 0 {
		query = fmt.Sprintf("%s limit=%d", query, limit)
	} else if limit == 0 && offset != 0 {
		query = fmt.Sprintf("%s offset=%d", query, offset)
	} else if limit != 0 && offset != 0 {
		query = fmt.Sprintf("%s limit=%d offset=%d", query, limit, offset)
	}
	return query
}

// QueryWhere ...
func QueryWhere(tableName string, conditions []*model.Condition) string {
	// build query
	query := fmt.Sprintf("select * from %s where", tableName)
	where := createQueriesWhere(conditions)
	query = fmt.Sprintf("%s%s", query, where)
	return query
}

// QueryWith ...
func QueryWith(tableName string, joins []*model.Join) string {
	// build query
	withs := createQueriesWith(joins)
	query := fmt.Sprintf("select * from %s %s", tableName, withs)
	return query
}

// QueryWhereWith ...
func QueryWhereWith(tableName string, joins []*model.Join, conditions []*model.Condition) string {

	// build query
	withs := createQueriesWith(joins)
	where := createQueriesWhere(conditions)
	query := fmt.Sprintf("select * from %s %s %s", tableName, withs, where)
	return query
}

// Create ...
func Create(tableName string, data interface{}) string {
	cols, values := createQueriesInsert(data)
	query := fmt.Sprintf("INSERT INTO %s %s %s", tableName, cols, values)
	return query
}

// Update ...
func Update(tableName string, data interface{}, conditions []*model.Condition) string {
	upd := createQueriesUpdate(data)
	query := fmt.Sprintf("UPDATE %s %s", tableName, upd)
	return query
}

// UpdateWhere ...
func UpdateWhere(tableName string, data interface{}, conditions []*model.Condition) string {
	upd := createQueriesUpdate(data)
	withs := createQueriesWhere(conditions)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, upd, withs)
	return query
}

// Delete ...
func Delete(tableName string, conditions []*model.Condition) string {
	withs := createQueriesWhere(conditions)
	query := fmt.Sprintf("DELETE FROM %s %s", tableName, withs)
	return query
}
