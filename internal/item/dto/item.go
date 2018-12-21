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
