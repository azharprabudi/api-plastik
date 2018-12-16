package command

import (
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// SellerCommand ...
type SellerCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
