package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// Supplier ...
type Supplier struct {
	SupplierID uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Phone      string    `db:"phone"`
	Address    string    `db:"address"`
	CreatedAt  time.Time `db:"created_at"`
}

// SupplierRead ...
type SupplierRead struct {
	Supplier
}

// SupplierCreate ...
type SupplierCreate struct {
	Supplier
}

// SupplierUpdate ...
type SupplierUpdate struct {
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Address string `db:"address"`
}
