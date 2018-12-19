package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ItemUnit ...
type ItemUnit struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// ItemUnitRes ...
type ItemUnitRes struct {
	ItemUnit
}
