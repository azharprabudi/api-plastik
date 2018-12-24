package query

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"
)

// TransactionQueryInterface ...
type TransactionQueryInterface interface {
	GetTransactions(uuid.UUID, int, int, string, string, string) ([]*model.TransactionRead, error)
	GetTransactionByID(uuid.UUID, uuid.UUID) (*model.TransactionReadDetail, error)
	GetTransactionEtcTypes(uuid.UUID) ([]*model.TransactionEtcTypeRead, error)
	GetTransactionEtcTypeByID(uuid.UUID, uuid.UUID) (*model.TransactionEtcTypeRead, error)
}
