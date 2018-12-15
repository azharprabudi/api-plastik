package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/api-plastik/internal/seller/dto"
	"github.com/api-plastik/internal/seller/model"
)

// TransformCreate ...
func (st *SellerTransform) TransformCreate(req *dto.SellerReq) *model.SellerCreate {
	create := &model.SellerCreate{
		Seller: model.Seller{
			SellerID:  uuid.NewV4(),
			Name:      req.Name,
			Address:   req.Address,
			Phone:     req.Phone,
			CreatedAt: time.Now().UTC(),
		},
	}
	return create
}

// TransformUpdate ...
func (st *SellerTransform) TransformUpdate(req *dto.SellerReq) *model.SellerUpdate {
	updt := &model.SellerUpdate{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
	return updt
}

// TransformGet ...
func (st *SellerTransform) TransformGet(s []*model.SellerRead) []*dto.SellerRes {
	// init variable
	var res = []*dto.SellerRes{}

	// transform data as dto expected
	for _, sel := range s {
		res = append(res, &dto.SellerRes{
			ID:        sel.Seller.SellerID,
			CreatedAt: sel.Seller.CreatedAt,
			SellerReq: dto.SellerReq{
				Name:    sel.Seller.Name,
				Phone:   sel.Seller.Phone,
				Address: sel.Seller.Address,
			},
		})
	}

	return res
}

// TransformGetByID ...
func (st *SellerTransform) TransformGetByID(s *model.SellerRead) *dto.SellerRes {
	return &dto.SellerRes{
		ID:        s.Seller.SellerID,
		CreatedAt: s.Seller.CreatedAt,
		SellerReq: dto.SellerReq{
			Name:    s.Seller.Name,
			Phone:   s.Seller.Phone,
			Address: s.Seller.Address,
		},
	}
}

// NewSellerTransform ...
func NewSellerTransform() SellerTransformInterface {
	return SellerTransformSingleton
}
