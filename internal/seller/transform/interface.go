package transform

import (
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	"github.com/azharprabudi/api-plastik/internal/seller/model"
	uuid "github.com/satori/go.uuid"
)

// SellerTransformInterface ...
type SellerTransformInterface interface {
	MakeModelCreateSeller(uuid.UUID, *dto.SellerReq) *model.SellerCreate
	MakeModelUpdateSeller(*dto.SellerReq) *model.SellerUpdate
	MakeResponseGetSellers([]*model.SellerRead) []*dto.SellerRes
	MakeResponseGetSellerByID(*model.SellerRead) *dto.SellerRes
}
