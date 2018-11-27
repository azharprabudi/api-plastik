package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// SellerID ...
type SellerID struct {
	SellerID uuid.UUID `db:"id"`
}

// SellerModelRead ...
type SellerModelRead struct {
	SellerModelCreate
}

// SellerModelCreate ...
type SellerModelCreate struct {
	SellerID
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}
