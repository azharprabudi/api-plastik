package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// ItemReq ...
type ItemReq struct {
	Name           string    `json:"name"`
	ItemCategoryID uuid.UUID `json:"itemCategoryId"`
	UnitID         uuid.UUID `json:"unitId"`
}

// ItemRes ...
type ItemRes struct {
	ItemID    uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ItemReq
}
