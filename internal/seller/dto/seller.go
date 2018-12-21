package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Seller ...
type Seller struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// SellerReq ...
type SellerReq struct {
	Seller
}

// SellerRes ...
type SellerRes struct {
	Seller
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
