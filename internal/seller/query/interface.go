package query

import (
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// SellerQueryInterface ...
type SellerQueryInterface interface {
	GetSellers(uuid.UUID) ([]*model.SellerRead, error)
	GetSellerByID(uuid.UUID, uuid.UUID) (*model.SellerRead, error)
}
