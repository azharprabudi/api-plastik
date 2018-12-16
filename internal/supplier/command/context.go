package command

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// SupplierCommand ...
type SupplierCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
