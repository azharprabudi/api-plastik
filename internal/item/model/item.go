package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Item ...
type Item struct {
	ItemID         uuid.UUID `db:"id"`
	ItemCategoryID uuid.UUID `db:"category_id"`
	Name           string    `db:"name"`
	UnitID         uuid.UUID `db:"unit_id"`
	CreatedAt      time.Time `db:"created_at"`
	CompanyID      uuid.UUID `db:"company_id"`
}

// ItemRead ...
type ItemRead struct {
	Item
	Active bool `db:"active"`
}

// ItemCreate ...
type ItemCreate struct {
	Item
}

// ItemUpdate ...
type ItemUpdate struct {
	Name           string    `db:"name"`
	ItemCategoryID uuid.UUID `db:"category_id"`
	UnitID         uuid.UUID `db:"unit_id"`
}

// ItemStockLogCreate ...
type ItemStockLogCreate struct {
	ID            uuid.UUID `db:"id"`
	ItemName      *string   `db:"item_name"`
	ItemID        uuid.UUID `db:"item_id"`
	TransactionID uuid.UUID `db:"transaction_id"`
	Qty           int       `db:"qty"`
	CreatedAt     time.Time `db:"created_at"`
}

// ItemStockLogRead ...
type ItemStockLogRead struct {
	ItemID   uuid.UUID `db:"item_id"`
	ItemName string    `db:"item_name"`
	Qty      int       `db:"qty"`
	UnitName string    `db:"unit_name"`
}
