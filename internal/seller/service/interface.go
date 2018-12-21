package service

import (
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	uuid "github.com/satori/go.uuid"
)

// SellerServiceInterface ...
type SellerServiceInterface interface {
	CreateSeller(*dto.SellerReq) (uuid.UUID, error)
	UpdateSeller(uuid.UUID, *dto.SellerReq) error
	DeleteSeller(uuid.UUID) error
	GetSellers() ([]*dto.SellerRes, error)
	GetSellerByID(uuid.UUID) (*dto.SellerRes, error)
}
