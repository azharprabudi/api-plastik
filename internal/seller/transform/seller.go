package transform

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	"github.com/azharprabudi/api-plastik/internal/seller/model"
)

// MakeModelCreateSeller ...
func (st *SellerTransform) MakeModelCreateSeller(companyID uuid.UUID, req *dto.SellerReq) *model.SellerCreate {
	return &model.SellerCreate{
		Seller: model.Seller{
			SellerID:  uuid.NewV4(),
			Name:      req.Name,
			Address:   req.Address,
			Phone:     req.Phone,
			CreatedAt: time.Now().UTC(),
			CompanyID: companyID,
		},
	}
}

// MakeModelUpdateSeller ...
func (st *SellerTransform) MakeModelUpdateSeller(req *dto.SellerReq) *model.SellerUpdate {
	return &model.SellerUpdate{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
}

// MakeResponseGetSellers ...
func (st *SellerTransform) MakeResponseGetSellers(res []*model.SellerRead) []*dto.SellerRes {
	var results = []*dto.SellerRes{}
	for _, s := range res {
		results = append(results, &dto.SellerRes{
			ID:        s.Seller.SellerID,
			CreatedAt: s.Seller.CreatedAt,
			Seller: dto.Seller{
				Name:    s.Seller.Name,
				Phone:   s.Seller.Phone,
				Address: s.Seller.Address,
			},
		})
	}
	return results
}

// MakeResponseGetSellerByID ...
func (st *SellerTransform) MakeResponseGetSellerByID(res *model.SellerRead) *dto.SellerRes {
	return &dto.SellerRes{
		ID:        res.Seller.SellerID,
		CreatedAt: res.Seller.CreatedAt,
		Seller: dto.Seller{
			Name:    res.Seller.Name,
			Phone:   res.Seller.Phone,
			Address: res.Seller.Address,
		},
	}
}

// NewSellerTransform ...
func NewSellerTransform() SellerTransformInterface {
	return SellerTransformSingleton
}
