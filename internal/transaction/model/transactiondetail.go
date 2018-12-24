package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TransactionDetail ...
type TransactionDetail struct {
	ID            uuid.UUID `db:"id"`
	TransactionID uuid.UUID `db:"transaction_id"`
	ItemID        uuid.UUID `db:"item_id"`
	ItemName      *string   `db:"item_name"`
	Amount        float64   `db:"amount"`
	Qty           int       `db:"qty"`
	CreatedAt     time.Time `db:"created_at"`
}

// TransactionDetailCreate ...
type TransactionDetailCreate struct {
	TransactionDetail
}

// TransactionDetailRead ...
type TransactionDetailRead struct {
	ID       *uuid.UUID `db:"id"`
	ItemID   *uuid.UUID `db:"item_id"`
	ItemName *string    `db:"item_name"`
	Amount   *float64   `db:"amount"`
	Qty      *int       `db:"qty"`
}
