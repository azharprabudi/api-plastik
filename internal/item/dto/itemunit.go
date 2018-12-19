package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ItemUnit ...
type ItemUnit struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// ItemUnitRes ...
type ItemUnitRes struct {
	ItemUnit
}
