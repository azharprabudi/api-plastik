package query

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// SellerQuery ...
type SellerQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
