package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/api-plastik/internal/seller/dto"
	"github.com/api-plastik/internal/seller/model"
)

// TransformCreate ...
func (it *SellerTransform) TransformCreate(sellerDTO *dto.SellerReq) *model.SellerCreate {
	sellerCreate := &model.SellerCreate{
		SellerID:  uuid.NewV4(),
		Name:      sellerDTO.Name,
		Address:   sellerDTO.Address,
		Phone:     sellerDTO.Phone,
		CreatedAt: time.Now().UTC(),
	}
	return sellerCreate
}

// TransformUpdate ...
func (it *SellerTransform) TransformUpdate(sellerDTO *dto.SellerReq) *model.SellerUpdate {
	sellerUpdate := &model.SellerUpdate{
		Name:    sellerDTO.Name,
		Address: sellerDTO.Address,
		Phone:   sellerDTO.Phone,
	}
	return sellerUpdate
}

// TransformGet ...
func (it *SellerTransform) TransformGet(sellerRead []*model.SellerRead) []*dto.SellerRes {
	// init variable
	var sellerRes = []*dto.SellerRes{}

	// transform data as dto expected
	for _, seller := range sellerRead {
		sellerRes = append(sellerRes, &dto.SellerRes{
			ID:        seller.SellerCreate.SellerID,
			CreatedAt: seller.SellerCreate.CreatedAt,
			SellerReq: dto.SellerReq{
				Name:    seller.SellerCreate.Name,
				Phone:   seller.SellerCreate.Phone,
				Address: seller.SellerCreate.Address,
			},
		})
	}

	return sellerRes
}

// TransformGetByID ...
func (it *SellerTransform) TransformGetByID(sellerRead *model.SellerRead) *dto.SellerRes {
	return &dto.SellerRes{
		ID:        sellerRead.SellerID,
		CreatedAt: sellerRead.CreatedAt,
		SellerReq: dto.SellerReq{
			Name:    sellerRead.Name,
			Phone:   sellerRead.Phone,
			Address: sellerRead.Address,
		},
	}
}

// NewSellerTransform ...
func NewSellerTransform() SellerTransformInterface {
	return SellerTransformSingleton
}
