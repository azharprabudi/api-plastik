package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction ...
type Transaction struct {
	ID         uuid.UUID  `db:"id"`
	Type       string     `db:"type"`
	Note       string     `db:"note"`
	UserID     uuid.UUID  `db:"user_id"`
	Amount     float64    `db:"amount"`
	SellerID   *uuid.UUID `db:"seller_id"`
	SupplierID *uuid.UUID `db:"supplier_id"`
	CreatedAt  time.Time  `db:"created_at"`
}

// TransactionRead ...
type TransactionRead struct {
	Transaction
	SellerName   *string `db:"seller_name"`
	SupplierName *string `db:"supplier_name"`
	UserName     *string `db:"user_name"`
}

// TransactionCreate ...
type TransactionCreate struct {
	Transaction
}

// TransactionReadDetail
type TransactionReadDetail struct {
	TransactionRead
	Details []*TransactionDetailRead
	Images  []*TransactionImageRead
}
