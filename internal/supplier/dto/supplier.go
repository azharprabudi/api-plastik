package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Supplier ...
type Supplier struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

// SupplierReq ...
type SupplierReq struct {
	Supplier
}

// SupplierRes ...
type SupplierRes struct {
	Supplier
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
