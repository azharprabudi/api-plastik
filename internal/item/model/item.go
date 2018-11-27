package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// ItemID ...
type ItemID struct {
	ItemID uuid.UUID `db:"id"`
}

// ItemModelRead ...
type ItemModelRead struct {
	ItemModelCreate
}

// ItemModelCreate ...
type ItemModelCreate struct {
	ItemID
	Name           string    `db:"name"`
	CategoryItemID int       `db:"category_id"`
	CreatedAt      time.Time `db:"created_at"`
}
