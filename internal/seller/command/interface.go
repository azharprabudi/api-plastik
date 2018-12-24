package command

import (
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// SellerCommandInterface ...
type SellerCommandInterface interface {
	CreateSeller(*model.SellerCreate) error
	UpdateSeller(uuid.UUID, uuid.UUID, *model.SellerUpdate) error
	DeleteSeller(uuid.UUID, uuid.UUID) error
}
