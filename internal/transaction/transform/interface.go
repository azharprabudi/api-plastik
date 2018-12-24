package transform

import (
	itemmodel "github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"
)

// TransactionTransformInterface ...
type TransactionTransformInterface interface {
	MakeModelCreateTransactionIn(uuid.UUID, *dto.TransactionInReq, string, uuid.UUID) *model.TransactionInCreate
	MakeModelCreateTransactionOut(uuid.UUID, *dto.TransactionOutReq, string, uuid.UUID) *model.TransactionOutCreate
	MakeModelCreateTransactionEtc(uuid.UUID, *dto.TransactionEtcReq, string, uuid.UUID) *model.TransactionEtcCreate
	MakeModelCreateTransactionDetails([]*dto.TransactionDetailReq, uuid.UUID) []*model.TransactionDetailCreate
	MakeModelCreateTransactionImages([]string, uuid.UUID) []*model.TransactionImageCreate
	MakeResponseGetTransactions([]*model.TransactionRead) []*dto.TransactionRes
	MakeResponseGetTransactionByID(*model.TransactionReadDetail) *dto.TransactionResDetail
	MakeModelCreateItemStockLog([]*dto.TransactionDetailReq, uuid.UUID, string) []*itemmodel.ItemStockLogCreate
	MakeResponseGetTransactionEtcTypes([]*model.TransactionEtcTypeRead) []*dto.TransactionEtcTypeRes
	MakeResponseGetTransactionEtcTypeByID(*model.TransactionEtcTypeRead) *dto.TransactionEtcTypeRes
	MakeModelCreateTransactionEtcType(uuid.UUID, *dto.TransactionEtcTypeReq) *model.TransactionEtcTypeCreate
	MakeModelUpdateTransactionEtcType(*dto.TransactionEtcTypeReq) *model.TransactionEtcTypeUpdate
}
