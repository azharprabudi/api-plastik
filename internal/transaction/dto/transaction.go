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
	ID     uuid.UUID `json:"id"`
	Amount float64   `json:"amount"`
	Transaction
	TransactionType string    `json:"type"`
	CreatedAt       time.Time `json:"createdAt"`
	SellerName      *string   `json:"seller_name"`
	SupplierName    *string   `json:"supplier_name"`
	UserName        *string   `json:"user_name"`
}

// TransactionResDetail ...
type TransactionResDetail struct {
	TransactionRes
	Details []*TransactionDetailRes
	Images  []*TransactionImageRes
}
