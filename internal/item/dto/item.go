package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Item struct {
	Name           string    `json:"name"`
	ItemCategoryID uuid.UUID `json:"itemCategoryId"`
	UnitID         uuid.UUID `json:"unitId"`
}

// ItemReq ...
type ItemReq struct {
	Item
}

// ItemRes ...
type ItemRes struct {
	ItemID    uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Item
}

// ItemStockLogRes ...
type ItemStockLogRes struct {
	ItemID   uuid.UUID `json:"id"`
	ItemName string    `json:"name"`
	Qty      int       `json:"qty"`
	UnitName string    `json:"unitName"`
}
