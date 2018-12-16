package transform

import (
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	"github.com/azharprabudi/api-plastik/internal/seller/model"
)

// SellerTransformInterface ...
type SellerTransformInterface interface {
	TransformCreate(*dto.SellerReq) *model.SellerCreate
	TransformUpdate(*dto.SellerReq) *model.SellerUpdate
	TransformGet([]*model.SellerRead) []*dto.SellerRes
	TransformGetByID(*model.SellerRead) *dto.SellerRes
}
