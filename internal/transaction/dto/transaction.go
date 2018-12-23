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
	Details []*TransactionDetailRes
	Images  []*TransactionImageRes
}
