package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// Seller ...
type Seller struct {
	SellerID  uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
}

// SellerRead ...
type SellerRead struct {
	Seller
}

// SellerCreate ...
type SellerCreate struct {
	Seller
}

// SellerUpdate ...
type SellerUpdate struct {
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Address string `db:"address"`
}
