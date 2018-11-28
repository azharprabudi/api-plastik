package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// SellerID ...
type SellerID struct {
	SellerID uuid.UUID `db:"id"`
}

// SellerRead ...
type SellerRead struct {
	SellerCreate
}

// SellerCreate ...
type SellerCreate struct {
	SellerID  uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}

// SellerUpdate ...
type SellerUpdate struct {
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Address string `db:"address"`
}
