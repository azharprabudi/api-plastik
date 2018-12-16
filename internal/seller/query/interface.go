package query

import (
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	"github.com/satori/go.uuid"
)

// SellerQueryInterface ...
type SellerQueryInterface interface {
	Get() ([]*model.SellerRead, error)
	GetByID(uuid.UUID) *model.SellerRead
}
