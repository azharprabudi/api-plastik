package dto

import uuid "github.com/satori/go.uuid"

// TransactionImageRes ...
type TransactionImageRes struct {
	ID    *uuid.UUID `json:"id"`
	Image *string    `json:"image"`
}
