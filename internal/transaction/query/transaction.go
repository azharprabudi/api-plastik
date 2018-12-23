package query

import (
	"fmt"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"

	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
)

// GetTransactions ...
func (tq *TransactionQuery) GetTransactions(limit int, start int, startAt string, endAt string, orderBy string) ([]*model.TransactionRead, error) {
	var queryLimit string
	var queryFilter string
	queryOrder := "order by transactions.created_at desc"
	var results []*model.TransactionRead

	if limit > 0 && start > 0 {
		queryLimit = fmt.Sprintf("limit %d offset %d", limit, start)
	}

	if startAt != "" && endAt != "" {
		queryFilter = fmt.Sprintf("where transactions.created_at::timestamp between '%s'::timestamp AND '%s'::timestamp", startAt, endAt)
	}

	if orderBy != "" {
		queryOrder = fmt.Sprintf("order by transactions.%s", orderBy)
	}

	query := fmt.Sprintf("select * , case when transactions.type = 'TRANSACTION_IN' then 'Transaksi Masuk' when transactions.type = 'TRANSACTION_OUT' then 'Transaksi Keluar' else 'Transaksi Lainnya' end as type_name from transactions %s %s %s", queryFilter, queryOrder, queryLimit)
	rows, err := tq.db.PgSQL.Queryx(query)
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

// GetTransactionEtcTypes ...
func (tq *TransactionQuery) GetTransactionEtcTypes() ([]*model.TransactionEtcTypeRead, error) {
	var results []*model.TransactionEtcTypeRead
	query := tq.qb.Query("transaction_etc_types", 0, 0, []*qbModel.Order{
		&qbModel.Order{
			Key:   "created_at",
			Value: "desc",
		},
	})
	rows, err := tq.db.PgSQL.Queryx(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tmp := new(model.TransactionEtcTypeRead)
		rows.StructScan(tmp)
		results = append(results, tmp)
	}

	return results, nil
}

// GetTransactionEtcTypeByID ...
func (tq *TransactionQuery) GetTransactionEtcTypeByID(id uuid.UUID) (*model.TransactionEtcTypeRead, error) {
	result := new(model.TransactionEtcTypeRead)
	query := tq.qb.QueryWhere("transaction_etc_types", []*qbModel.Condition{&qbModel.Condition{
		Key:      "id",
		NextCond: "",
		Operator: "=",
		Value:    id.String(),
	}}, nil)
	err := tq.db.PgSQL.QueryRowx(query).StructScan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// NewTransactionQuery ...
func NewTransactionQuery(db *db.DB) TransactionQueryInterface {
	q := qb.NewQueryBuilder()
	return &TransactionQuery{
		qb: q,
		db: db,
	}

}
