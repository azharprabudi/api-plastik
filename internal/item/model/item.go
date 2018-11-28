package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// ItemID ...
type ItemID struct {
	ItemID uuid.UUID `db:"id"`
}

// ItemRead ...
type ItemRead struct {
	ItemCreate
}

// ItemCreate ...
type ItemCreate struct {
	ItemID         uuid.UUID `db:"id"`
	CategoryItemID int       `db:"category_id"`
	Name           string    `db:"name"`
	CreatedAt      time.Time `db:"created_at"`
}

// ItemUpdate ...
type ItemUpdate struct {
	Name           string `db:"name"`
	CategoryItemID int    `db:"category_id"`
}
