package query

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// ExpenseQuery ...
type ExpenseQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
