package command

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// ItemCommand ...
type ItemCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
