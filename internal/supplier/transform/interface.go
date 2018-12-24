package transform

import (
	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
	uuid "github.com/satori/go.uuid"
)

// SupplierTransformInterface ...
type SupplierTransformInterface interface {
	MakeModelCreateSupplier(uuid.UUID, *dto.SupplierReq) *model.SupplierCreate
	MakeModelUpdateSupplier(*dto.SupplierReq) *model.SupplierUpdate
	MakeResponseGetSuppliers([]*model.SupplierRead) []*dto.SupplierRes
	MakeResponseGetSupplierByID(*model.SupplierRead) *dto.SupplierRes
}
