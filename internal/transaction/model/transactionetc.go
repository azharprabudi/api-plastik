package model

import uuid "github.com/satori/go.uuid"

// TransactionEtc ...
type TransactionEtc struct {
	ID                 uuid.UUID `db:"id"`
	TransactionEtcType uuid.UUID `db:"transaction_etc_type"`
}

// TransactionEtcCreate ...
type TransactionEtcCreate struct {
	TransactionEtc
	TransactionCreate
}
