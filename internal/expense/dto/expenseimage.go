package dto

import (
	"github.com/satori/go.uuid"
)

// ExpenseImageRes ...
type ExpenseImageRes struct {
	ExpenseImageID uuid.UUID `json:"id"`
	Image          string    `json:"image"`
}
