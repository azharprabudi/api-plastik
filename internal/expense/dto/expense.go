package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// ExpenseReq ...
type ExpenseReq struct {
	Name          string   `json:"name"`
	ExpenseTypeID int      `json:"expense_type_id"`
	Amount        float64  `json:"amount"`
	Note          string   `json:"note"`
	Images        []string `json:"images"`
}

// ExpenseRes ...
type ExpenseRes struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ExpenseReq
}
