package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Seller ...
type Seller struct {
	SellerID  uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
	CompanyID uuid.UUID `db:"company_id"`
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
