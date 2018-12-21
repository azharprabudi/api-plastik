package query

import (
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// SellerQueryInterface ...
type SellerQueryInterface interface {
	GetSellers() ([]*model.SellerRead, error)
	GetSellerByID(uuid.UUID) (*model.SellerRead, error)
}
