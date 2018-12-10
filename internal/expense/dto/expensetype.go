package dto

import "time"

// ExpenseTypeReq ...
type ExpenseTypeReq struct {
	Name string `json:"name"`
}

// ExpenseTypeRes ...
type ExpenseTypeRes struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
