package model

import uuid "github.com/satori/go.uuid"

// TransactionOut ...
type TransactionOut struct {
	ID       uuid.UUID `db:"id"`
	SellerID uuid.UUID `db:"seller_id"`
}

// TransactionOutCreate ...
type TransactionOutCreate struct {
	TransactionOut
	TransactionCreate
}
