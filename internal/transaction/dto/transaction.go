package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction ...
type Transaction struct {
	Type       string     `json:"type"`
	Note       string     `json:"note"`
	SellerID   *uuid.UUID `json:"sellerId"`
	SupplierID *uuid.UUID `json:"supplierId"`
}

// TransactionReq ...
type TransactionReq struct {
	Transaction
	Details []TransactionDetailReq `json:"details"`
	Images  []string               `json:"images"`
}

// TransactionRes ....
type TransactionRes struct {
	ID     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
	Transaction
	CreatedAt    time.Time `json:"createdAt"`
	SellerName   *string   `json:"seller_name"`
	SupplierName *string   `json:"supplier_name"`
	UserName     *string   `json:"user_name"`
}

// TransactionResDetail ...
type TransactionResDetail struct {
	TransactionRes
	Details []*TransactionDetailRes
	Images  []*TransactionImageRes
}
