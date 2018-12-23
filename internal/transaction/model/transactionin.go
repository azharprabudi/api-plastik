package model

import uuid "github.com/satori/go.uuid"

// TransactionIn ...
type TransactionIn struct {
	ID         uuid.UUID `db:"id"`
	SupplierID uuid.UUID `db:"supplier_id"`
}

// TransactionInCreate ...
type TransactionInCreate struct {
	TransactionIn
	TransactionCreate
}
