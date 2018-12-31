package service

import (
	"github.com/azharprabudi/api-plastik/db"
	trx "github.com/azharprabudi/api-plastik/helper/transaction"
	"github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/azharprabudi/api-plastik/internal/transaction/command"
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	"github.com/azharprabudi/api-plastik/internal/transaction/event"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	"github.com/azharprabudi/api-plastik/internal/transaction/query"
	"github.com/azharprabudi/api-plastik/internal/transaction/transform"
	"github.com/azharprabudi/api-plastik/internal/user/value"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// FindTransactions ...
func (ts *TransactionService) FindTransactions(companyID uuid.UUID, page int, startAt string, endAt string, orderBy string) ([]*dto.TransactionRes, error) {
	limit := 10
	start := (page * limit) - limit
	transactions, err := ts.query.GetTransactions(companyID, limit, start, startAt, endAt, orderBy)
	if err != nil {
		return nil, err
	}

	return ts.transform.MakeResponseGetTransactions(transactions), nil
}

// CreateTransactionIn ...
func (ts *TransactionService) CreateTransactionIn(companyID uuid.UUID, req *dto.TransactionInReq, transactionType string) (uuid.UUID, error) {
	transaction := ts.transform.MakeModelCreateTransactionIn(companyID, req, transactionType, value.USER_ID)
	details := ts.transform.MakeModelCreateTransactionDetails(req.Details, transaction.Transaction.ID)
	detailsStock := ts.transform.MakeModelCreateItemStockLog(req.Details, transaction.Transaction.ID, transactionType)
	images := ts.transform.MakeModelCreateTransactionImages(req.Images, transaction.Transaction.ID)

	trans := trx.NewTransaction()
	ctx := trans.CreateTrx(ts.db.PgSQL)
	err := trans.RunTrx(ctx, func(tx *sqlx.Tx) error {
		err := ts.command.CreateTransaction(tx, &(*transaction).TransactionCreate)
		if err != nil {
			return err
		}

		err = ts.command.CreateTransactionIn(tx, &(*transaction).TransactionIn, transaction.Transaction.ID)
		if err != nil {
			return err
		}

		err = ts.createTransacionDetails(tx, details, companyID)
		if err != nil {
			return err
		}

		err = ts.createTransacionImages(tx, images)
		if err != nil {
			return err
		}

		err = ts.event.TriggerAfterCreateTransaction(tx, detailsStock, companyID)
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

// CreateTransactionOut ...
func (ts *TransactionService) CreateTransactionOut(companyID uuid.UUID, req *dto.TransactionOutReq, transactionType string) (uuid.UUID, error) {
	transaction := ts.transform.MakeModelCreateTransactionOut(companyID, req, transactionType, value.USER_ID)
	details := ts.transform.MakeModelCreateTransactionDetails(req.Details, transaction.Transaction.ID)
	detailsStock := ts.transform.MakeModelCreateItemStockLog(req.Details, transaction.Transaction.ID, transactionType)
	images := ts.transform.MakeModelCreateTransactionImages(req.Images, transaction.Transaction.ID)

	trans := trx.NewTransaction()
	ctx := trans.CreateTrx(ts.db.PgSQL)
	err := trans.RunTrx(ctx, func(tx *sqlx.Tx) error {
		err := ts.command.CreateTransaction(tx, &(*transaction).TransactionCreate)
		if err != nil {
			return err
		}

		err = ts.command.CreateTransactionOut(tx, &(*transaction).TransactionOut, transaction.Transaction.ID)
		if err != nil {
			return err
		}

		err = ts.createTransacionDetails(tx, details, companyID)
		if err != nil {
			return err
		}

		err = ts.createTransacionImages(tx, images)
		if err != nil {
			return err
		}

		err = ts.event.TriggerAfterCreateTransaction(tx, detailsStock, companyID)
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

// CreateTransactionEtc ...
func (ts *TransactionService) CreateTransactionEtc(companyID uuid.UUID, req *dto.TransactionEtcReq, transactionType string) (uuid.UUID, error) {
	transaction := ts.transform.MakeModelCreateTransactionEtc(companyID, req, transactionType, value.USER_ID)
	images := ts.transform.MakeModelCreateTransactionImages(req.Images, transaction.Transaction.ID)

	trans := trx.NewTransaction()
	ctx := trans.CreateTrx(ts.db.PgSQL)
	err := trans.RunTrx(ctx, func(tx *sqlx.Tx) error {
		err := ts.command.CreateTransaction(tx, &(*transaction).TransactionCreate)
		if err != nil {
			return err
		}

		err = ts.command.CreateTransactionEtc(tx, &(*transaction).TransactionEtc, transaction.Transaction.ID)
		if err != nil {
			return err
		}

		err = ts.createTransacionImages(tx, images)
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

// createTransacionDetails ...
func (ts *TransactionService) createTransacionDetails(tx *sqlx.Tx, details []*model.TransactionDetailCreate, companyID uuid.UUID) error {
	var err error
	for _, detail := range details {
		item, _err := ts.itemService.GetItemByID(companyID, detail.TransactionDetail.ItemID)
		if _err != nil {
			err = _err
			break
		}

		detail.TransactionDetail.ItemName = &(*item).Item.Name
		_err = ts.command.CreateTransactionDetail(tx, detail)
		if _err != nil {
			err = _err
			break
		}
	}

	if err != nil {
		return err
	}
	return nil
}

// createTransacionImages ...
func (ts *TransactionService) createTransacionImages(tx *sqlx.Tx, images []*model.TransactionImageCreate) error {
	var err error
	for _, image := range images {
		_err := ts.command.CreateTransactionImage(tx, image)
		if _err != nil {
			err = _err
			break
		}
	}

	if err != nil {
		return err
	}
	return nil
}

// FindTransactionByID ...
func (ts *TransactionService) FindTransactionByID(companyID uuid.UUID, id uuid.UUID) (*dto.TransactionResDetail, error) {
	transaction, err := ts.query.GetTransactionByID(companyID, id)
	if err != nil {
		return nil, err
	}

	return ts.transform.MakeResponseGetTransactionByID(transaction), nil
}

// FindTransactionEtcTypes ...
func (ts *TransactionService) FindTransactionEtcTypes(companyID uuid.UUID) ([]*dto.TransactionEtcTypeRes, error) {
	transactionEtcTypes, err := ts.query.GetTransactionEtcTypes(companyID)
	if err != nil {
		return nil, err
	}

	return ts.transform.MakeResponseGetTransactionEtcTypes(transactionEtcTypes), nil
}

// FindTransactionEtcTypeByID ...
func (ts *TransactionService) FindTransactionEtcTypeByID(companyID uuid.UUID, id uuid.UUID) (*dto.TransactionEtcTypeRes, error) {
	transactionEtcType, err := ts.query.GetTransactionEtcTypeByID(companyID, id)
	if err != nil {
		return nil, err
	}

	return ts.transform.MakeResponseGetTransactionEtcTypeByID(transactionEtcType), nil
}

// CreateTransactionEtcType ...
func (ts *TransactionService) CreateTransactionEtcType(companyID uuid.UUID, req *dto.TransactionEtcTypeReq) (uuid.UUID, error) {
	transactionEtcType := ts.transform.MakeModelCreateTransactionEtcType(companyID, req)
	err := ts.command.CreateTransactionEtcType(transactionEtcType)
	if err != nil {
		return uuid.Nil, err
	}

	return transactionEtcType.TransactionEtcType.ID, nil
}

// UpdateTransactionEtcType ...
func (ts *TransactionService) UpdateTransactionEtcType(companyID uuid.UUID, id uuid.UUID, req *dto.TransactionEtcTypeReq) error {
	transactionEtcType := ts.transform.MakeModelUpdateTransactionEtcType(req)
	err := ts.command.UpdateTransactionEtcType(companyID, id, transactionEtcType)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTransactionEtcType ...
func (ts *TransactionService) DeleteTransactionEtcType(companyID uuid.UUID, id uuid.UUID) error {
	err := ts.command.DeleteTransactionEtcType(companyID, id)
	if err != nil {
		return err
	}
	return nil
}

// NewTransactionService ...
func NewTransactionService(db *db.DB) TransactionServiceInterface {
	itemService := service.NewItemService(db)
	return &TransactionService{
		db:          db,
		query:       query.NewTransactionQuery(db),
		command:     command.NewTransactionCommand(db),
		transform:   transform.NewTransactionTransform(),
		itemService: itemService,
		event:       event.NewTransactionEvent(itemService),
	}
}
