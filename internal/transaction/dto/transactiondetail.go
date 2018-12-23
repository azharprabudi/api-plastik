package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TransactionDetail ...
type TransactionDetail struct {
	ItemID uuid.UUID `json:"itemId"`
	Qty    int       `json:"qty"`
	Amount float64   `json:"amount"`
}

// TransactionDetailReq ...
type TransactionDetailReq struct {
	ItemID uuid.UUID `json:"itemId"`
	Qty    int       `json:"qty"`
	Amount float64   `json:"amount"`
}

// TransactionDetailRes ...
type TransactionDetailRes struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transactionId"`
	TransactionDetail
	CreatedAt time.Time `json:"createdAt"`
}
