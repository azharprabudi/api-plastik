package dto

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TransactionEtcType ...
type TransactionEtcType struct {
	Name string `json:"name"`
}

// TransactionEtcTypeReq ...
type TransactionEtcTypeReq struct {
	TransactionEtcType
}

// TransactionEtcTypeRes ...
type TransactionEtcTypeRes struct {
	TransactionEtcType
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
