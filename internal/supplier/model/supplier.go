package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// SupplierID ...
type SupplierID struct {
	SupplierID uuid.UUID `db:"supplierId"`
}

// SupplierModelRead ...
type SupplierModelRead struct {
	SupplierModelCreate
}

// SupplierModelCreate ...
type SupplierModelCreate struct {
	SupplierID
	Name      string    `db:"supplierName"`
	Phone     *string   `db:"supplierPhone"`   // nullable
	Address   *int      `db:"supplierAddress"` // nullable
	CreatedAt time.Time `db:"createdAt"`
}
