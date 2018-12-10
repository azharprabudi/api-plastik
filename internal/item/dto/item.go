package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// ItemReq ...
type ItemReq struct {
	Name           string `json:"name"`
	ItemCategoryID int    `json:"itemCategoryId"`
	UnitID         int    `json:"unitId"`
}

// ItemRes ...
type ItemRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ItemReq
}
