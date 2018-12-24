package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction ...
type Transaction struct {
	ID        uuid.UUID `db:"id"`
	Type      string    `db:"type"`
	Note      string    `db:"note"`
	UserID    uuid.UUID `db:"user_id"`
	Amount    float64   `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
	CompanyID uuid.UUID `db:"company_id"`
}

// TransactionRead ...
type TransactionRead struct {
	Transaction
	TypeName string `db:"type_name"`
}

// TransactionCreate ...
type TransactionCreate struct {
	Transaction
}

// TransactionReadDetail ...
type TransactionReadDetail struct {
	TransactionRead
	TransactionOutID       *uuid.UUID `db:"transaction_out_id"`
	SellerID               *uuid.UUID `db:"seller_id"`
	SellerName             *string    `db:"seller_name"`
	TransactionInID        *uuid.UUID `db:"transaction_in_id"`
	SupplierID             *uuid.UUID `db:"supplier_id"`
	SupplierName           *string    `db:"supplier_name"`
	TransactionEtcID       *uuid.UUID `db:"transaction_etc_id"`
	TransactionEtcTypeName *string    `db:"transaction_etc_type_name"`
	Details                []*TransactionDetailRead
	Images                 []*TransactionImageRead
}
