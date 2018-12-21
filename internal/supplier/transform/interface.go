package transform

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
)

// SupplierTransformInterface ...
type SupplierTransformInterface interface {
	MakeModelCreateSupplier(*dto.SupplierReq) *model.SupplierCreate
	MakeModelUpdateSupplier(*dto.SupplierReq) *model.SupplierUpdate
	MakeResponseGetSuppliers([]*model.SupplierRead) []*dto.SupplierRes
	MakeResponseGetSupplierByID(*model.SupplierRead) *dto.SupplierRes
}
