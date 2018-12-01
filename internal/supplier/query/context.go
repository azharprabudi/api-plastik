package query

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// SupplierQuery ...
type SupplierQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
