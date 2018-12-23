package dto

import uuid "github.com/satori/go.uuid"

// TransactionIn ...
type TransactionIn struct {
	SupplierID uuid.UUID `json:"supplierId"`
}

// TransactionInReq ...
type TransactionInReq struct {
	Transaction
	TransactionIn
	Details []*TransactionDetailReq `json:"details"`
	Images  []string                `json:"images"`
}

// TransactionInRes ...
type TransactionInRes struct {
	Transaction
	TransactionIn
	Details []*TransactionDetailRes `json:"details"`
	Images  []*TransactionImageRes  `json:"images"`
}
