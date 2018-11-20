package helpers

import "github.com/jmoiron/sqlx"

// DBQueryBuilder ...
type DBQueryBuilder struct {
	db        *sqlx.DB
	tableName string
	columns   []string
	limit     int
	offset    int
}
