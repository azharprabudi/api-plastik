package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction ...
type Transaction struct {
	Note string `json:"note"`
}

// TransactionRes ....
type TransactionRes struct {
	Transaction
	ID                  uuid.UUID `json:"id"`
	Amount              float64   `json:"amount"`
	TransactionType     string    `json:"type"`
	TransactionTypeName string    `json:"typeName"`
	CreatedAt           time.Time `json:"createdAt"`
}

// TransactionResDetail ...
type TransactionResDetail struct {
	TransactionRes
	TransactionOutID       *uuid.UUID              `json:"transactionOutId"`
	SellerID               *uuid.UUID              `json:"sellerId"`
	SellerName             *string                 `json:"sellerName"`
	TransactionInID        *uuid.UUID              `json:"transactionInId"`
	SupplierID             *uuid.UUID              `json:"supplierId"`
	SupplierName           *string                 `json:"supplierName"`
	TransactionEtcID       *uuid.UUID              `json:"transactionEtcId"`
	TransactionEtcTypeName *string                 `json:"transactionEtcTypeName"`
	Details                []*TransactionDetailRes `json:"details"`
	Images                 []*TransactionImageRes  `json:"images"`
}
