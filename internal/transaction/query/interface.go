package query

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"
)

// TransactionCommandInterface ...
type TransactionQueryInterface interface {
	GetTransactions(int, int, string, string, string) ([]*model.TransactionRead, error)
	GetTransactionByID(uuid.UUID) (*model.TransactionReadDetail, error)
}
