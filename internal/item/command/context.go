package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// ItemCommand ...
type ItemCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
