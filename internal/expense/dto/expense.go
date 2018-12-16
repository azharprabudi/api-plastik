package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// ExpenseReq ...
type ExpenseReq struct {
	ExpenseTypeID uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	Note          string    `json:"note"`
	Images        []string  `json:"images"`
}

// ExpenseRes ...
type ExpenseRes struct {
	ExpenseID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ExpenseReq
}
