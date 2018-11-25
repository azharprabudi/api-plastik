package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// ItemQuery ...
type ItemQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
