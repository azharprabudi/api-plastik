package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ItemCategory ...
type ItemCategory struct {
	ItemCategoryID uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	CreatedAt      time.Time `db:"created_at"`
	CompanyID      uuid.UUID `db:"company_id"`
}

// ItemCategoryRead ...
type ItemCategoryRead struct {
	ItemCategory
	Active bool `db:"active"`
}

// ItemCategoryCreate ...
type ItemCategoryCreate struct {
	ItemCategory
}

// ItemCategoryUpdate ...
type ItemCategoryUpdate struct {
	Name string `db:"name"`
}
