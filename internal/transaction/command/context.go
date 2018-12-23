package command

import (
	"github.com/azharprabudi/api-plastik/db"

	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// TransactionCommand ...
type TransactionCommand struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
