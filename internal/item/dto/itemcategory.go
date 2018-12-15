package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// ItemCategoryReq ...
type ItemCategoryReq struct {
	Name string `json:"name"`
}

// ItemCategoryRes ...
type ItemCategoryRes struct {
	ItemCategoryReq
	ItemCategoryID uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
}
