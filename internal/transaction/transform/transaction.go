package transform

import (
	"time"

	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"
)

// MakeModelCreateTransaction ...
func (tt *TransactionTransform) MakeModelCreateTransaction(req *dto.TransactionReq, userID uuid.UUID) *model.TransactionCreate {
	var totalAmount float64
	for _, detail := range req.Details {
		totalAmount += detail.Amount
	}

	return &model.TransactionCreate{
		Transaction: model.Transaction{
			ID:         uuid.NewV4(),
			Note:       req.Transaction.Note,
			SellerID:   req.Transaction.SellerID,
			SupplierID: req.Transaction.SupplierID,
			Type:       req.Transaction.Type,
			UserID:     userID,
			Amount:     totalAmount,
			CreatedAt:  time.Now().UTC(),
		},
	}
}

// MakeModelCreateTransactionDetails ...
func (tt *TransactionTransform) MakeModelCreateTransactionDetails(req *dto.TransactionReq, id uuid.UUID) []*model.TransactionDetailCreate {
	var results []*model.TransactionDetailCreate
	for _, detail := range req.Details {
		results = append(results, &model.TransactionDetailCreate{
			TransactionDetail: model.TransactionDetail{
				ID:            uuid.NewV4(),
				TransactionID: id,
				Amount:        detail.Amount,
				CreatedAt:     time.Now().UTC(),
				Qty:           detail.Qty,
				ItemID:        detail.ItemID,
				ItemName:      nil,
			},
		})
	}

	return results
}

// MakeModelCreateTransactionImages ...
func (tt *TransactionTransform) MakeModelCreateTransactionImages(req *dto.TransactionReq, id uuid.UUID) []*model.TransactionImageCreate {
	var results []*model.TransactionImageCreate
	for _, image := range req.Images {
		results = append(results, &model.TransactionImageCreate{
			TransactionImage: model.TransactionImage{
				ID:            uuid.NewV4(),
				TransactionID: id,
				Image:         image,
				CreatedAt:     time.Now().UTC(),
			},
		})
	}

	return results
}

// MakeResponseGetTransaction ...
func (tt *TransactionTransform) MakeResponseGetTransaction(transactions []*model.TransactionRead) []*dto.TransactionRes {
	var results []*dto.TransactionRes
	for _, trx := range transactions {
		results = append(results, &dto.TransactionRes{
			ID:     trx.Transaction.ID,
			Amount: trx.Transaction.Amount,
			Transaction: dto.Transaction{
				Note:       trx.Transaction.Note,
				SellerID:   trx.Transaction.SellerID,
				SupplierID: trx.Transaction.SupplierID,
				Type:       trx.Transaction.Type,
			},
			CreatedAt: trx.Transaction.CreatedAt,
		})
	}
	return results
}

// MakeResponseGetTransactionByID ...
func (tt *TransactionTransform) MakeResponseGetTransactionByID(transaction *model.TransactionReadDetail) *dto.TransactionResDetail {
	return nil
}

func NewTransactionTransform() TransactionTransformInterface {
	return &TransactionTransform{}
}
