package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction ...
type Transaction struct {
	ID        uuid.UUID `db:"id"`
	Type      string    `db:"type"`
	Note      string    `db:"note"`
	CompanyID uuid.UUID `db:"company_id"`
	UserID    uuid.UUID `db:"user_id"`
	Amount    float64   `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}

// TransactionRead ...
type TransactionRead struct {
	Transaction
	TypeName string `db:"type_name"`
}

// TransactionCreate ...
type TransactionCreate struct {
	Transaction
}

// TransactionReadDetail ...
type TransactionReadDetail struct {
	TransactionRead
	Details []*TransactionDetailRead
	Images  []*TransactionImageRead
}
