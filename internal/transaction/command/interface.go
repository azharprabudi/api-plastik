package command

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// TransactionCommandInterface ...
type TransactionCommandInterface interface {
	CreateTransaction(*sqlx.Tx, *model.TransactionCreate) error
	CreateTransactionIn(*sqlx.Tx, *model.TransactionIn, uuid.UUID) error
	CreateTransactionOut(*sqlx.Tx, *model.TransactionOut, uuid.UUID) error
	CreateTransactionEtc(*sqlx.Tx, *model.TransactionEtc, uuid.UUID) error
	CreateTransactionDetail(*sqlx.Tx, *model.TransactionDetailCreate) error
	CreateTransactionImage(*sqlx.Tx, *model.TransactionImageCreate) error
	CreateTransactionEtcType(*model.TransactionEtcTypeCreate) error
	UpdateTransactionEtcType(uuid.UUID, *model.TransactionEtcTypeUpdate) error
	DeleteTransactionEtcType(uuid.UUID) error
}
