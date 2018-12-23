package dto

import uuid "github.com/satori/go.uuid"

// TransactionEtc ...
type TransactionEtc struct {
	Amount               float64   `json:"amount"`
	TransactionEtcTypeID uuid.UUID `json:"type"`
}

// TransactionEtcReq ...
type TransactionEtcReq struct {
	Transaction
	TransactionEtc
	Images []string `json:"images"`
}

// TransactionEtcRes ...
type TransactionEtcRes struct {
	Transaction
	TransactionEtc
	Images []*TransactionImageRes `json:"images"`
}
