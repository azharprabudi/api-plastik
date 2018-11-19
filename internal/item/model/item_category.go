package model

import (
	"time"
)

// ItemCategoryID ...
type ItemCategoryID struct {
	ItemCategoryID int `db:"itemCategoryId"`
}

// ItemCategoryModelRead ...
type ItemCategoryModelRead struct {
	ItemCategoryID
	ItemCategoryCreate
}

// ItemCategoryCreate ...
type ItemCategoryCreate struct {
	Name      string    `db:"itemName"`
	CreatedAt time.Time `db:"createdAt"`
}
