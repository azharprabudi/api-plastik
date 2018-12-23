package dto

import uuid "github.com/satori/go.uuid"

// TransactionOut ...
type TransactionOut struct {
	SellerID uuid.UUID `json:"sellerId"`
}

// TransactionOutReq ...
type TransactionOutReq struct {
	Transaction
	TransactionOut
	Images  []string                `json:"images"`
	Details []*TransactionDetailReq `json:"details"`
}

// TransactionOutRes ...
type TransactionOutRes struct {
	Transaction
	TransactionOut
	Images  []*TransactionImageRes  `json:"images"`
	Details []*TransactionDetailRes `json:"details"`
}
