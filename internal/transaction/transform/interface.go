package transform

import (
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"
)

// TransactionTransformInterface ...
type TransactionTransformInterface interface {
	// command
	MakeModelCreateTransaction(*dto.TransactionReq, uuid.UUID) *model.TransactionCreate
	MakeModelCreateTransactionDetails(*dto.TransactionReq, uuid.UUID) []*model.TransactionDetailCreate
	MakeModelCreateTransactionImages(*dto.TransactionReq, uuid.UUID) []*model.TransactionImageCreate

	// query
	MakeResponseGetTransaction([]*model.TransactionRead) []*dto.TransactionRes
	MakeResponseGetTransactionByID(*model.TransactionReadDetail) *dto.TransactionResDetail
}
