package transform

import (
	"github.com/api-plastik/internal/supplier/dto"
	"github.com/api-plastik/internal/supplier/model"
)

// SupplierTransformInterface ...
type SupplierTransformInterface interface {
	TransformCreate(*dto.SupplierReq) *model.SupplierCreate
	TransformUpdate(*dto.SupplierReq) *model.SupplierUpdate
	TransformGet([]*model.SupplierRead) []*dto.SupplierRes
	TransformGetByID(*model.SupplierRead) *dto.SupplierRes
}
