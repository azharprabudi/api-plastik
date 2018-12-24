package command

import (
	"fmt"

	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbModel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// CreateTransaction ...
func (tc *TransactionCommand) CreateTransaction(tx *sqlx.Tx, transaction *model.TransactionCreate) error {
	query := fmt.Sprintf("INSERT INTO transactions (\"id\", \"note\", \"user_id\", \"company_id\", \"type\", \"amount\", \"created_at\") VALUES ($1, $2, $3, $4, $5, $6, $7)")
	_, err := tx.Exec(query, transaction.Transaction.ID, transaction.Transaction.Note, transaction.Transaction.UserID, transaction.Transaction.CompanyID, transaction.Transaction.Type, transaction.Transaction.Amount, transaction.Transaction.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionIn ...
func (tc *TransactionCommand) CreateTransactionIn(tx *sqlx.Tx, in *model.TransactionIn, id uuid.UUID) error {
	query := fmt.Sprintf("INSERT INTO transactions_in (\"id\", \"supplier_id\", \"transaction_id\") VALUES ($1, $2, $3)")
	_, err := tx.Exec(query, in.ID, in.SupplierID, id)
	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionOut ...
func (tc *TransactionCommand) CreateTransactionOut(tx *sqlx.Tx, out *model.TransactionOut, id uuid.UUID) error {
	query := fmt.Sprintf("INSERT INTO transactions_out (\"id\", \"seller_id\", \"transaction_id\") VALUES ($1, $2, $3)")
	_, err := tx.Exec(query, out.ID, out.SellerID, id)
	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionEtc ...
func (tc *TransactionCommand) CreateTransactionEtc(tx *sqlx.Tx, etc *model.TransactionEtc, id uuid.UUID) error {
	query := fmt.Sprintf("INSERT INTO transactions_etc (\"id\", \"transaction_etc_type\", \"transaction_id\") VALUES ($1, $2, $3)")
	_, err := tx.Exec(query, etc.ID, etc.TransactionEtcType, id)
	if err != nil {
		return err
	}

	return nil
}

// CreateTransactionDetail ...
func (tc *TransactionCommand) CreateTransactionDetail(tx *sqlx.Tx, transactionDetail *model.TransactionDetailCreate) error {
	query := tc.qb.Create("transaction_details", (*transactionDetail).TransactionDetail)
	_, err := tx.Exec(query, transactionDetail.TransactionDetail.ID, transactionDetail.TransactionDetail.TransactionID, transactionDetail.TransactionDetail.ItemID, transactionDetail.TransactionDetail.ItemName, transactionDetail.TransactionDetail.Amount, transactionDetail.TransactionDetail.Qty, transactionDetail.TransactionDetail.CreatedAt)
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

// CreateTransactionEtcType ...
func (tc *TransactionCommand) CreateTransactionEtcType(transactionEtcType *model.TransactionEtcTypeCreate) error {
	query := tc.qb.Create("transaction_etc_types", (*transactionEtcType).TransactionEtcType)
	_, err := tc.db.PgSQL.Exec(query, transactionEtcType.TransactionEtcType.ID, transactionEtcType.TransactionEtcType.Name, transactionEtcType.TransactionEtcType.CreatedAt, transactionEtcType.TransactionEtcType.CompanyID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTransactionEtcType ...
func (tc *TransactionCommand) UpdateTransactionEtcType(companyID uuid.UUID, id uuid.UUID, transactionEtcType *model.TransactionEtcTypeUpdate) error {
	query := tc.qb.UpdateWhere("transaction_etc_types", *transactionEtcType, []*qbModel.Condition{&qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbModel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := tc.db.PgSQL.Exec(query, transactionEtcType.Name)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTransactionEtcType ...
func (tc *TransactionCommand) DeleteTransactionEtcType(companyID uuid.UUID, id uuid.UUID) error {
	query := tc.qb.Delete("transaction_etc_types", []*qbModel.Condition{&qbModel.Condition{
		Key:      "id",
		Operator: "=",
		NextCond: "AND",
		Value:    id.String(),
	}, &qbModel.Condition{
		Key:      "company_id",
		Operator: "=",
		NextCond: "",
		Value:    companyID.String(),
	}})
	_, err := tc.db.PgSQL.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// NewTransactionCommand ...
func NewTransactionCommand(db *db.DB) TransactionCommandInterface {
	q := qb.NewQueryBuilder()
	return &TransactionCommand{
		db: db,
		qb: q,
	}
}
