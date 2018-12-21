package service

import (
	"github.com/azharprabudi/api-plastik/internal/seller/transform"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/seller/command"
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	"github.com/azharprabudi/api-plastik/internal/seller/query"
)

// GetSeller ...
func (ss *SellerService) GetSellers() ([]*dto.SellerRes, error) {
	sellers, err := ss.query.GetSellers()
	if err != nil {
		return nil, err
	}

	return ss.transform.MakeResponseGetSellers(sellers), nil
}

// GetSellerByID ...
func (ss *SellerService) GetSellerByID(id uuid.UUID) (*dto.SellerRes, error) {
	seller, err := ss.query.GetSellerByID(id)
	if err != nil {
		return nil, err
	}

	return ss.transform.MakeResponseGetSellerByID(seller), nil
}

// CreateSeller ...
func (ss *SellerService) CreateSeller(req *dto.SellerReq) (uuid.UUID, error) {
	seller := ss.transform.MakeModelCreateSeller(req)
	err := ss.command.CreateSeller(seller)
	if err != nil {
		return uuid.Nil, err
	}

	return seller.Seller.SellerID, nil
}

// UpdateSeller ...
func (ss *SellerService) UpdateSeller(id uuid.UUID, req *dto.SellerReq) error {
	seller := ss.transform.MakeModelUpdateSeller(req)
	err := ss.command.UpdateSeller(id, seller)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSeller ...
func (ss *SellerService) DeleteSeller(id uuid.UUID) error {
	err := ss.command.DeleteSeller(id)
	if err != nil {
		return err
	}

	return nil
}

// NewSellerService ...
func NewSellerService(db *db.DB) SellerServiceInterface {
	return &SellerService{
		query:     query.NewSellerQuery(db),
		command:   command.NewSellerCommand(db),
		transform: transform.NewSellerTransform(),
	}
}
