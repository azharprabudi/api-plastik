package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// ItemID ...
type ItemID struct {
	ItemID uuid.UUID `db:"itemId"`
}

// ItemModelRead ...
type ItemModelRead struct {
	ItemModelCreate
}

// ItemModelCreate ...
type ItemModelCreate struct {
	ItemID
	Name           string    `db:"itemName"`
	CategoryItemID int       `db:"itemCategoryId"`
	CreatedAt      time.Time `db:"createdAt"`
}
