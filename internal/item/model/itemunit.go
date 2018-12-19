package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type ItemUnit struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type ItemUnitRead struct {
	ItemUnit
}
