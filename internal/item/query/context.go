package query

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// ItemQuery ...
type ItemQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
