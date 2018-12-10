package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// ExpenseCommand ...
type ExpenseCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}