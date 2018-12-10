package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// ExpenseQuery ...
type ExpenseQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
