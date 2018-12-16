package query

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// SupplierQuery ...
type SupplierQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
