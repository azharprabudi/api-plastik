package dto

import "time"

// ItemCategoryReq ...
type ItemCategoryReq struct {
	Name string `json:"name"`
}

// ItemCategoryRes ...
type ItemCategoryRes struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
