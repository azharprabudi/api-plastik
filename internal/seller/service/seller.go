package service

import (
	"github.com/azharprabudi/api-plastik/internal/seller/transform"
	"github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/seller/command"
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	"github.com/azharprabudi/api-plastik/internal/seller/query"
)

// GetSeller ...
func (sellerService *SellerService) GetSeller() ([]*dto.SellerRes, error) {
	// add data to db
	sellerModel, err := sellerService.query.Get()
	if err != nil {
		return nil, err
	}

	// transform data from model
	sellerDTO := sellerService.transform.TransformGet(sellerModel)
	return sellerDTO, nil
}

// GetSellerByID ...
func (sellerService *SellerService) GetSellerByID(sellerID uuid.UUID) *dto.SellerRes {
	sellerModel := sellerService.query.GetByID(sellerID)
	if sellerModel == nil {
		return nil
	}

	// transform data from model
	sellerDTO := sellerService.transform.TransformGetByID(sellerModel)
	return sellerDTO
}

// CreateSeller ...
func (sellerService *SellerService) CreateSeller(item *dto.SellerReq) (uuid.UUID, error) {
	// transform dto to model
	sellerCreate := sellerService.transform.TransformCreate(item)

	// add data to db
	err := sellerService.command.Create(sellerCreate)
	if err != nil {
		return uuid.Nil, err
	}
	return sellerCreate.SellerID, nil
}

// UpdateSeller ...
func (sellerService *SellerService) UpdateSeller(sellerID uuid.UUID, item *dto.SellerReq) error {
	// transform dto to model
	sellerUpdate := sellerService.transform.TransformUpdate(item)

	// update to db
	err := sellerService.command.Update(sellerID, sellerUpdate)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSeller ...
func (sellerService *SellerService) DeleteSeller(sellerID uuid.UUID) error {
	// delete data from db
	err := sellerService.command.Delete(sellerID)
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
