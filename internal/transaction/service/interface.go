package service

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	uuid "github.com/satori/go.uuid"
)

// TransactionServiceInterface ...
type TransactionServiceInterface interface {
	FindTransactions(int, string, string, string) ([]*dto.TransactionRes, error)
	CreateTransactionIn(*dto.TransactionInReq, string) (uuid.UUID, error)
	CreateTransactionOut(*dto.TransactionOutReq, string) (uuid.UUID, error)
	CreateTransactionEtc(*dto.TransactionEtcReq, string) (uuid.UUID, error)
	FindTransactionByID(uuid.UUID) (*dto.TransactionResDetail, error)
	FindTransactionEtcTypes() ([]*dto.TransactionEtcTypeRes, error)
	FindTransactionEtcTypeByID(uuid.UUID) (*dto.TransactionEtcTypeRes, error)
	CreateTransactionEtcType(*dto.TransactionEtcTypeReq) (uuid.UUID, error)
	UpdateTransactionEtcType(uuid.UUID, *dto.TransactionEtcTypeReq) error
	DeleteTransactionEtcType(uuid.UUID) error
}
