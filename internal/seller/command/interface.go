package command

import (
	"github.com/api-plastik/internal/seller/model"
	"github.com/satori/go.uuid"
)

// SellerCommandInterface ...
type SellerCommandInterface interface {
	// item
	Create(*model.SellerCreate) error
	Update(uuid.UUID, *model.SellerUpdate) error
	Delete(uuid.UUID) error
}
