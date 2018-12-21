package command

import (
	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	"github.com/jmoiron/sqlx"
)

// CreateTransaction ...
func (tc *TransactionCommand) CreateTransaction(tx *sqlx.Tx, transaction *model.TransactionCreate) error {
	query := tc.qb.Create("transactions", (*transaction).Transaction)
	_, err := tx.Exec(query, transaction.ID, transaction.Type, transaction.Note, transaction.UserID, transaction.Amount, transaction.SellerID, transaction.SupplierID, transaction.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionDetail ...
func (tc *TransactionCommand) CreateTransactionDetail(tx *sqlx.Tx, transactionDetail *model.TransactionDetailCreate) error {
	query := tc.qb.Create("transaction_details", (*transactionDetail).TransactionDetail)
	_, err := tx.Exec(query, transactionDetail.TransactionDetail.ID, transactionDetail.TransactionDetail.TransactionID, transactionDetail.TransactionDetail.Amount, transactionDetail.TransactionDetail.Qty, transactionDetail.TransactionDetail.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionImage ...
func (tc *TransactionCommand) CreateTransactionImage(tx *sqlx.Tx, transactionImage *model.TransactionImageCreate) error {
	query := tc.qb.Create("transaction_images", (*transactionImage).TransactionImage)
	_, err := tx.Exec(query, transactionImage.TransactionImage.ID, transactionImage.TransactionImage.TransactionID, transactionImage.TransactionImage.Image, transactionImage.TransactionImage.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionCommand(db *db.DB) TransactionCommandInterface {
	q := qb.NewQueryBuilder()
	return &TransactionCommand{
		db: db,
		qb: q,
	}
}
