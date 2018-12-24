package service

import (
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	uuid "github.com/satori/go.uuid"
)

// SellerServiceInterface ...
type SellerServiceInterface interface {
	CreateSeller(uuid.UUID, *dto.SellerReq) (uuid.UUID, error)
	UpdateSeller(uuid.UUID, uuid.UUID, *dto.SellerReq) error
	DeleteSeller(uuid.UUID, uuid.UUID) error
	GetSellers(uuid.UUID) ([]*dto.SellerRes, error)
	GetSellerByID(uuid.UUID, uuid.UUID) (*dto.SellerRes, error)
}
