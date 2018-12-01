package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// SupplierCommand ...
type SupplierCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
