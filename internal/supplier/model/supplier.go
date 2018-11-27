package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// SupplierID ...
type SupplierID struct {
	SupplierID uuid.UUID `db:"id"`
}

// SupplierModelRead ...
type SupplierModelRead struct {
	SupplierModelCreate
}

// SupplierModelCreate ...
type SupplierModelCreate struct {
	SupplierID
	Name      string    `db:"name"`
	Phone     *string   `db:"phone"`   // nullable
	Address   *int      `db:"address"` // nullable
	CreatedAt time.Time `db:"created_at"`
}
