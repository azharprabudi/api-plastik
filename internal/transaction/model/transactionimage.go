package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TransactionImage ...
type TransactionImage struct {
	ID            uuid.UUID `db:"id"`
	TransactionID uuid.UUID `db:"transaction_id"`
	Image         string    `db:"image"`
	CreatedAt     time.Time `db:"created_at"`
}

// TransactionImageCreate ...
type TransactionImageCreate struct {
	TransactionImage
}

// TransactionImageRead ...
type TransactionImageRead struct {
	ID    *uuid.UUID `db:"id"`
	Image *string    `db:"image"`
}
