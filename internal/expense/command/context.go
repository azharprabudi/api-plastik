package command

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// ExpenseCommand ...
type ExpenseCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
