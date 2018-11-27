package model

import (
	"time"
)

// ItemCategoryID ...
type ItemCategoryID struct {
	ItemCategoryID int `db:"id"`
}

// ItemCategoryModelRead ...
type ItemCategoryModelRead struct {
	ItemCategoryID
	ItemCategoryCreate
}

// ItemCategoryCreate ...
type ItemCategoryCreate struct {
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
