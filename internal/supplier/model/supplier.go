package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Supplier ...
type Supplier struct {
	SupplierID uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Phone      string    `db:"phone"`
	Address    string    `db:"address"`
	CreatedAt  time.Time `db:"created_at"`
	CompanyID  uuid.UUID `db:"company_id"`
}

// SupplierRead ...
type SupplierRead struct {
	Supplier
	Active bool `db:"active"`
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
