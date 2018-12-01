package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// SupplierID ...
type SupplierID struct {
	SupplierID uuid.UUID `db:"id"`
}

// SupplierRead ...
type SupplierRead struct {
	SupplierCreate
}

// SupplierCreate ...
type SupplierCreate struct {
	SupplierID uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Phone      string    `db:"phone"`
	Address    string    `db:"address"`
	CreatedAt  time.Time `db:"created_at"`
}

// SupplierUpdate ...
type SupplierUpdate struct {
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Address string `db:"address"`
}
