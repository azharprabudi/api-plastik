package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// SellerID ...
type SellerID struct {
	SellerID uuid.UUID `db:"sellerId"`
}

// SellerModelRead ...
type SellerModelRead struct {
	SellerModelCreate
}

// SellerModelCreate ...
type SellerModelCreate struct {
	SellerID
	Name      string    `db:"sellerName"`
	Phone     string    `db:"sellerphone"`
	Address   string    `db:"sellerAddress"`
	CreatedAt time.Time `db:"createdAt"`
}
