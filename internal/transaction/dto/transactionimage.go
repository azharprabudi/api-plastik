package dto

import uuid "github.com/satori/go.uuid"

// TransactionImageRes ...
type TransactionImageRes struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transactionId"`
	Image         string    `json:"image"`
}
