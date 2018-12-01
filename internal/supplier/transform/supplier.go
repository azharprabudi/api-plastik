package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/api-plastik/internal/supplier/dto"
	"github.com/api-plastik/internal/supplier/model"
)

// TransformCreate ...
func (it *SupplierTransform) TransformCreate(supplierDTO *dto.SupplierReq) *model.SupplierCreate {
	supplierCreate := &model.SupplierCreate{
		SupplierID: uuid.NewV4(),
		Name:       supplierDTO.Name,
		Address:    supplierDTO.Address,
		Phone:      supplierDTO.Phone,
		CreatedAt:  time.Now().UTC(),
	}
	return supplierCreate
}

// TransformUpdate ...
func (it *SupplierTransform) TransformUpdate(supplierDTO *dto.SupplierReq) *model.SupplierUpdate {
	supplierUpdate := &model.SupplierUpdate{
		Name:    supplierDTO.Name,
		Address: supplierDTO.Address,
		Phone:   supplierDTO.Phone,
	}
	return supplierUpdate
}

// TransformGet ...
func (it *SupplierTransform) TransformGet(supplierRead []*model.SupplierRead) []*dto.SupplierRes {
	// init variable
	var supplierRes = []*dto.SupplierRes{}

	// transform data as dto expected
	for _, supplier := range supplierRead {
		supplierRes = append(supplierRes, &dto.SupplierRes{
			ID:        supplier.SupplierCreate.SupplierID,
			CreatedAt: supplier.SupplierCreate.CreatedAt,
			SupplierReq: dto.SupplierReq{
				Name:    supplier.SupplierCreate.Name,
				Phone:   supplier.SupplierCreate.Phone,
				Address: supplier.SupplierCreate.Address,
			},
		})
	}

	return supplierRes
}

// TransformGetByID ...
func (it *SupplierTransform) TransformGetByID(supplierRead *model.SupplierRead) *dto.SupplierRes {
	return &dto.SupplierRes{
		ID:        supplierRead.SupplierID,
		CreatedAt: supplierRead.CreatedAt,
		SupplierReq: dto.SupplierReq{
			Name:    supplierRead.Name,
			Phone:   supplierRead.Phone,
			Address: supplierRead.Address,
		},
	}
}

// NewSupplierTransform ...
func NewSupplierTransform() SupplierTransformInterface {
	return SupplierTransformSingleton
}
