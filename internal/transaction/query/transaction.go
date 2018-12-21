package query

import (
	"fmt"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"

	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
)

// GetTransactions ...
func (tq *TransactionQuery) GetTransactions(limit int, start int, startAt string, endAt string, orderBy string) ([]*model.TransactionRead, error) {
	var queryLimit string
	var queryFilter string
	var queryOrder string = "ORDER BY transactions.created_at DESC"
	var results []*model.TransactionRead

	if limit > 0 && start > 0 {
		queryLimit = fmt.Sprintf("LIMIT %d OFFSET %d", limit, start)
	}

	if startAt != "" && endAt != "" {
		queryFilter = fmt.Sprintf("WHERE transactions.created_at BETWEEN %s AND %s", startAt, endAt)
	}

	if orderBy != "" {
		queryOrder = fmt.Sprintf("ORDER BY transactions.%s", orderBy)
	}

	rows, err := tq.db.PgSQL.Queryx(fmt.Sprintf("SELECT transactions.*, suppliers.name as supplier_name, sellers.name as seller_name FROM transactions LEFT JOIN sellers ON transactions.seller_id = sellers.id LEFT JOIN suppliers ON transactions.supplier_id = suppliers.id %s %s %s", queryLimit, queryFilter, queryOrder))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.TransactionRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetTransactionByID ...
func (tq *TransactionQuery) GetTransactionByID(id uuid.UUID) (*model.TransactionReadDetail, error) {
	// tq.db.PgSQL.Queryx(fmt.Sprintf("SELECT transactions.*, suppliers.name as supplier_name, sellers.name as seller_name FROM transactions LEFT JOIN sellers ON transactions.seller_id = sellers.id LEFT JOIN suppliers ON transactions.supplier_id = suppliers.id JOIN "))

	return nil, nil
}

func NewTransactionQuery(db *db.DB) TransactionQueryInterface {
	q := qb.NewQueryBuilder()
	return &TransactionQuery{
		qb: q,
		db: db,
	}

}
