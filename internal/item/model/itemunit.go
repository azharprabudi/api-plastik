package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ItemUnit ...
type ItemUnit struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

// ItemUnitRead ...
type ItemUnitRead struct {
	ItemUnit
}
