package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// SupplierReq ...
type SupplierReq struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// SupplierRes ...
type SupplierRes struct {
	ID uuid.UUID `json:"id"`
	SupplierReq
	CreatedAt time.Time `json:"createdAt"`
}
