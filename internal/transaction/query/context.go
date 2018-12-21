package query

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// TransactionQuery ...
type TransactionQuery struct {
	db *db.DB
	qb qb.QueryBuilderInterface
}
