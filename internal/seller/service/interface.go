package service

import (
	"github.com/api-plastik/internal/seller/dto"
	"github.com/satori/go.uuid"
)

// SellerServiceInterface ...
type SellerServiceInterface interface {
	CreateSeller(*dto.SellerReq) (uuid.UUID, error)
	UpdateSeller(uuid.UUID, *dto.SellerReq) error
	DeleteSeller(uuid.UUID) error
	GetSeller() ([]*dto.SellerRes, error)
	GetSellerByID(uuid.UUID) *dto.SellerRes
}
