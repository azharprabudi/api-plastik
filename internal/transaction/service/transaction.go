package service

import (
	"github.com/azharprabudi/api-plastik/db"
	trx "github.com/azharprabudi/api-plastik/helper/transaction"
	"github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/azharprabudi/api-plastik/internal/transaction/command"
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	"github.com/azharprabudi/api-plastik/internal/transaction/query"
	"github.com/azharprabudi/api-plastik/internal/transaction/transform"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// FindTransactions ...
func (ts *TransactionService) FindTransactions(page int, startAt string, endAt string, orderBy string) ([]*dto.TransactionRes, error) {
	limit := 10
	start := (page * limit) - (limit - 1)
	transactions, err := ts.query.GetTransactions(limit, start, startAt, endAt, orderBy)
	if err != nil {
		return nil, err
	}

	return ts.transform.MakeResponseGetTransaction(transactions), nil
}

// CreateTransaction ...
func (ts *TransactionService) CreateTransaction(req *dto.TransactionReq) (uuid.UUID, error) {
	transaction := ts.transform.MakeModelCreateTransaction(req, uuid.NewV4())
	details := ts.transform.MakeModelCreateTransactionDetails(req, transaction.ID)
	images := ts.transform.MakeModelCreateTransactionImages(req, transaction.ID)

	trans := trx.NewTransaction()
	ctx := trans.CreateTrx(ts.db.PgSQL)
	err := trans.RunTrx(ctx, func(tx *sqlx.Tx) error {
		err := ts.command.CreateTransaction(tx, transaction)
		if err != nil {
			return err
		}

		for _, detail := range details {
			item, err := ts.itemService.GetItemByID(detail.TransactionDetail.ItemID)
			if err != nil {
				return err
			}

			detail.TransactionDetail.ItemName = &(*item).Item.Name
			err = ts.command.CreateTransactionDetail(tx, detail)
			if err != nil {
				return err
			}
		}

		for _, image := range images {
			err = ts.command.CreateTransactionImage(tx, image)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return uuid.Nil, err
	}
	return transaction.Transaction.ID, nil
}

// FindTransactionByID ...
func (ts *TransactionService) FindTransactionByID(id uuid.UUID) (*dto.TransactionResDetail, error) {
	transaction, err := ts.query.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}

	return ts.transform.MakeResponseGetTransactionByID(transaction), nil
}

// NewSupplierService ...
func NewTransactionService(db *db.DB, itemService service.ItemServiceInterface) TransactionServiceInterface {
	return &TransactionService{
		db:          db,
		query:       query.NewTransactionQuery(db),
		command:     command.NewTransactionCommand(db),
		transform:   transform.NewTransactionTransform(),
		itemService: service.NewItemService(db),
	}
}
