package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionDetail struct {
	ItemID uuid.UUID `json:"itemId"`
	Qty    int       `json:"qty"`
	Amount float64   `json:"amount"`
}

type TransactionDetailReq struct {
	TransactionDetail
}

type TransactionDetailRes struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transactionId"`
	TransactionDetail
	CreatedAt time.Time `json:"createdAt"`
}
