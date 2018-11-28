package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// SellerReq ...
type SellerReq struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// SellerRes ...
type SellerRes struct {
	ID uuid.UUID `json:"id"`
	SellerReq
	CreatedAt time.Time `json:"createdAt"`
}
