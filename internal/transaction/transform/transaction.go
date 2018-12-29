package transform

import (
	"time"

	"github.com/azharprabudi/api-plastik/internal/transaction/value"

	itemModel "github.com/azharprabudi/api-plastik/internal/item/model"
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	"github.com/azharprabudi/api-plastik/internal/transaction/model"
	uuid "github.com/satori/go.uuid"
)

// MakeModelCreateTransactionIn ...
func (tt *TransactionTransform) MakeModelCreateTransactionIn(companyID uuid.UUID, req *dto.TransactionInReq, transactionType string, userID uuid.UUID) *model.TransactionInCreate {
	var totalAmount float64
	for _, detail := range req.Details {
		totalAmount += detail.Amount * float64(detail.Qty)
	}

	return &model.TransactionInCreate{
		TransactionCreate: model.TransactionCreate{
			Transaction: model.Transaction{
				ID:        uuid.NewV4(),
				Note:      req.Transaction.Note,
				Type:      transactionType,
				UserID:    userID,
				Amount:    totalAmount,
				CreatedAt: time.Now().UTC(),
				CompanyID: companyID,
			},
		},
		TransactionIn: model.TransactionIn{
			ID:         uuid.NewV4(),
			SupplierID: req.TransactionIn.SupplierID,
		},
	}
}

// MakeModelCreateTransactionOut ...
func (tt *TransactionTransform) MakeModelCreateTransactionOut(companyID uuid.UUID, req *dto.TransactionOutReq, transactionType string, userID uuid.UUID) *model.TransactionOutCreate {
	var totalAmount float64
	for _, detail := range req.Details {
		totalAmount += detail.Amount * float64(detail.Qty)
	}

	return &model.TransactionOutCreate{
		TransactionCreate: model.TransactionCreate{
			Transaction: model.Transaction{
				ID:        uuid.NewV4(),
				Note:      req.Transaction.Note,
				Type:      transactionType,
				UserID:    userID,
				Amount:    totalAmount,
				CreatedAt: time.Now().UTC(),
				CompanyID: companyID,
			},
		},
		TransactionOut: model.TransactionOut{
			ID:       uuid.NewV4(),
			SellerID: req.TransactionOut.SellerID,
		},
	}
}

// MakeModelCreateTransactionEtc ...
func (tt *TransactionTransform) MakeModelCreateTransactionEtc(companyID uuid.UUID, req *dto.TransactionEtcReq, transactionType string, userID uuid.UUID) *model.TransactionEtcCreate {
	return &model.TransactionEtcCreate{
		TransactionCreate: model.TransactionCreate{
			Transaction: model.Transaction{
				ID:        uuid.NewV4(),
				Note:      req.Transaction.Note,
				Type:      transactionType,
				UserID:    userID,
				CompanyID: companyID,
				Amount:    req.TransactionEtc.Amount,
				CreatedAt: time.Now().UTC(),
			},
		},
		TransactionEtc: model.TransactionEtc{
			ID:                 uuid.NewV4(),
			TransactionEtcType: req.TransactionEtc.TransactionEtcTypeID,
		},
	}
}

// MakeModelCreateTransactionDetails ...
func (tt *TransactionTransform) MakeModelCreateTransactionDetails(req []*dto.TransactionDetailReq, id uuid.UUID) []*model.TransactionDetailCreate {
	var results []*model.TransactionDetailCreate
	for _, detail := range req {
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
func (tt *TransactionTransform) MakeModelCreateTransactionImages(req []string, id uuid.UUID) []*model.TransactionImageCreate {
	var results []*model.TransactionImageCreate
	for _, image := range req {
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

// MakeResponseGetTransactions ...
func (tt *TransactionTransform) MakeResponseGetTransactions(req []*model.TransactionRead) []*dto.TransactionRes {
	var results = []*dto.TransactionRes{}
	for _, trx := range req {
		results = append(results, &dto.TransactionRes{
			ID:     trx.Transaction.ID,
			Amount: trx.Transaction.Amount,
			Transaction: dto.Transaction{
				Note: trx.Transaction.Note,
			},
			TransactionType:     trx.Transaction.Type,
			TransactionTypeName: trx.TypeName,
			CreatedAt:           trx.Transaction.CreatedAt,
		})
	}
	return results
}

// MakeModelCreateItemStockLog ...
func (tt *TransactionTransform) MakeModelCreateItemStockLog(req []*dto.TransactionDetailReq, id uuid.UUID, transactionType string) []*itemModel.ItemStockLogCreate {
	var results []*itemModel.ItemStockLogCreate
	for _, transaction := range req {
		var qty int
		if transactionType == value.TRANSACTION_IN {
			qty = transaction.Qty
		} else {
			qty = -1 * transaction.Qty
		}

		results = append(results, &itemModel.ItemStockLogCreate{
			ID:            uuid.NewV4(),
			ItemName:      nil,
			ItemID:        transaction.ItemID,
			Qty:           qty,
			TransactionID: id,
			CreatedAt:     time.Now().UTC(),
		})
	}

	return results
}

// MakeResponseGetTransactionEtcTypes ...
func (tt *TransactionTransform) MakeResponseGetTransactionEtcTypes(res []*model.TransactionEtcTypeRead) []*dto.TransactionEtcTypeRes {
	results := []*dto.TransactionEtcTypeRes{}
	for _, t := range res {
		results = append(results, &dto.TransactionEtcTypeRes{
			ID: t.TransactionEtcType.ID,
			TransactionEtcType: dto.TransactionEtcType{
				Name: t.TransactionEtcType.Name,
			},
			CreatedAt: t.TransactionEtcType.CreatedAt,
		})
	}

	return results
}

// MakeResponseGetTransactionEtcTypeByID ...
func (tt *TransactionTransform) MakeResponseGetTransactionEtcTypeByID(res *model.TransactionEtcTypeRead) *dto.TransactionEtcTypeRes {
	return &dto.TransactionEtcTypeRes{
		TransactionEtcType: dto.TransactionEtcType{
			Name: res.Name,
		},
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
	}
}

// MakeModelCreateTransactionEtcType ...
func (tt *TransactionTransform) MakeModelCreateTransactionEtcType(companyID uuid.UUID, req *dto.TransactionEtcTypeReq) *model.TransactionEtcTypeCreate {
	return &model.TransactionEtcTypeCreate{
		TransactionEtcType: model.TransactionEtcType{
			ID:        uuid.NewV4(),
			Name:      req.TransactionEtcType.Name,
			CreatedAt: time.Now().UTC(),
			CompanyID: companyID,
		},
	}
}

// MakeModelUpdateTransactionEtcType ...
func (tt *TransactionTransform) MakeModelUpdateTransactionEtcType(req *dto.TransactionEtcTypeReq) *model.TransactionEtcTypeUpdate {
	return &model.TransactionEtcTypeUpdate{
		Name: req.TransactionEtcType.Name,
	}
}

// MakeResponseGetTransactionByID ...
func (tt *TransactionTransform) MakeResponseGetTransactionByID(transaction *model.TransactionReadDetail) *dto.TransactionResDetail {
	images := []*dto.TransactionImageRes{}
	details := []*dto.TransactionDetailRes{}
	for _, image := range transaction.Images {
		images = append(images, &dto.TransactionImageRes{
			ID:    image.ID,
			Image: image.Image,
		})
	}

	for _, detail := range transaction.Details {
		details = append(details, &dto.TransactionDetailRes{
			ID:       detail.ID,
			ItemID:   detail.ItemID,
			Qty:      detail.Qty,
			Amount:   detail.Amount,
			ItemName: detail.ItemName,
		})
	}

	return &dto.TransactionResDetail{
		TransactionRes: dto.TransactionRes{
			Amount:    transaction.Amount,
			CreatedAt: transaction.CreatedAt,
			ID:        transaction.ID,
			Transaction: dto.Transaction{
				Note: transaction.Note,
			},
			TransactionType:     transaction.Type,
			TransactionTypeName: transaction.TypeName,
		},
		TransactionInID:        transaction.TransactionInID,
		SupplierID:             transaction.SupplierID,
		SupplierName:           transaction.SupplierName,
		TransactionOutID:       transaction.TransactionOutID,
		SellerID:               transaction.SellerID,
		SellerName:             transaction.SellerName,
		TransactionEtcID:       transaction.TransactionEtcID,
		TransactionEtcTypeName: transaction.TransactionEtcTypeName,
		Details:                details,
		Images:                 images,
	}
}

// NewTransactionTransform ...
func NewTransactionTransform() TransactionTransformInterface {
	return &TransactionTransform{}
}
