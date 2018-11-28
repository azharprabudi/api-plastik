package command

import (
	"github.com/api-plastik/db"
	"github.com/api-plastik/helper/querybuilder"
)

// SellerCommand ...
type SellerCommand struct {
	db *db.DB
	q  qb.QueryBuilderInterface
}
