package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// SellerQuery ...
type SellerQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
