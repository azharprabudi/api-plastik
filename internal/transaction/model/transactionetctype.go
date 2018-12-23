package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TransactionEtcType ...
type TransactionEtcType struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

// TransactionEtcTypeRead ...
type TransactionEtcTypeRead struct {
	TransactionEtcType
}

// TransactionEtcTypeCreate ...
type TransactionEtcTypeCreate struct {
	TransactionEtcType
}

// TransactionEtcTypeUpdate ...
type TransactionEtcTypeUpdate struct {
	Name string `db:"name"`
}
