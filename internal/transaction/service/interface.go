package service

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	uuid "github.com/satori/go.uuid"
)

// TransactionServiceInterface ...
type TransactionServiceInterface interface {
	FindTransactions(uuid.UUID, int, string, string, string) ([]*dto.TransactionRes, error)
	CreateTransactionIn(uuid.UUID, *dto.TransactionInReq, string) (uuid.UUID, error)
	CreateTransactionOut(uuid.UUID, *dto.TransactionOutReq, string) (uuid.UUID, error)
	CreateTransactionEtc(uuid.UUID, *dto.TransactionEtcReq, string) (uuid.UUID, error)
	FindTransactionByID(uuid.UUID, uuid.UUID) (*dto.TransactionResDetail, error)
	FindTransactionEtcTypes(uuid.UUID) ([]*dto.TransactionEtcTypeRes, error)
	FindTransactionEtcTypeByID(uuid.UUID, uuid.UUID) (*dto.TransactionEtcTypeRes, error)
	CreateTransactionEtcType(uuid.UUID, *dto.TransactionEtcTypeReq) (uuid.UUID, error)
	UpdateTransactionEtcType(uuid.UUID, uuid.UUID, *dto.TransactionEtcTypeReq) error
	DeleteTransactionEtcType(uuid.UUID, uuid.UUID) error
	GetCountTransactions(uuid.UUID, string, string) (int, error)
	GetSummaryTransactions(uuid.UUID, string, string) (float64, error)
	GetSummaryTransactionsIn(uuid.UUID, string, string) (float64, error)
	GetSummaryTransactionsOut(uuid.UUID, string, string) (float64, error)
	GetSummaryTransactionsEtc(uuid.UUID, string, string) (float64, error)
}
