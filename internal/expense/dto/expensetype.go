package dto

import (
	"time"

	"github.com/satori/go.uuid"
)

// ExpenseTypeReq ...
type ExpenseTypeReq struct {
	Name string `json:"name"`
}

// ExpenseTypeRes ...
type ExpenseTypeRes struct {
	ExpenseTypeReq
	ExpenseTypeID uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
}
