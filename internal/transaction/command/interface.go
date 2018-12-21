package command

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	"github.com/jmoiron/sqlx"
)

// TransactionCommandInterface ...
type TransactionCommandInterface interface {
	CreateTransaction(*sqlx.Tx, *model.TransactionCreate) error
	CreateTransactionDetail(*sqlx.Tx, *model.TransactionDetailCreate) error
	CreateTransactionImage(*sqlx.Tx, *model.TransactionImageCreate) error
}
