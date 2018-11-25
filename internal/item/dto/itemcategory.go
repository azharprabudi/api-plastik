package dto

import "time"

// ItemCategoryIncReq ...
type ItemCategoryIncReq struct {
	Name string `json:"name"`
}

// ItemCategoryIncRes ...
type ItemCategoryIncRes struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
