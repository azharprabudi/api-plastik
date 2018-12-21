package service

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	uuid "github.com/satori/go.uuid"
)

// TransactionServiceInterface ...
type TransactionServiceInterface interface {
	FindTransactions(int, string, string, string) ([]*dto.TransactionRes, error)
	CreateTransaction(*dto.TransactionReq) (uuid.UUID, error)
	FindTransactionByID(uuid.UUID) (*dto.TransactionResDetail, error)
}
