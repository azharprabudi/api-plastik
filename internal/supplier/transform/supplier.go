package transform

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
)

// TransformCreate ...
func (st *SupplierTransform) TransformCreate(s *dto.SupplierReq) *model.SupplierCreate {
	create := &model.SupplierCreate{
		Supplier: model.Supplier{
			SupplierID: uuid.NewV4(),
			Name:       s.Name,
			Address:    s.Address,
			Phone:      s.Phone,
			CreatedAt:  time.Now().UTC(),
		},
	}
	return create
}

// TransformUpdate ...
func (st *SupplierTransform) TransformUpdate(s *dto.SupplierReq) *model.SupplierUpdate {
	supplierUpdate := &model.SupplierUpdate{
		Name:    s.Name,
		Address: s.Address,
		Phone:   s.Phone,
	}
	return supplierUpdate
}

// TransformGet ...
func (st *SupplierTransform) TransformGet(s []*model.SupplierRead) []*dto.SupplierRes {
	// init variable
	var res = []*dto.SupplierRes{}

	// transform data as dto expected
	for _, supplier := range s {
		res = append(res, &dto.SupplierRes{
			ID:        supplier.Supplier.SupplierID,
			CreatedAt: supplier.Supplier.CreatedAt,
			SupplierReq: dto.SupplierReq{
				Name:    supplier.Supplier.Name,
				Phone:   supplier.Supplier.Phone,
				Address: supplier.Supplier.Address,
			},
		})
	}

	return res
}

// TransformGetByID ...
func (st *SupplierTransform) TransformGetByID(s *model.SupplierRead) *dto.SupplierRes {
	return &dto.SupplierRes{
		ID:        s.SupplierID,
		CreatedAt: s.CreatedAt,
		SupplierReq: dto.SupplierReq{
			Name:    s.Name,
			Phone:   s.Phone,
			Address: s.Address,
		},
	}
}

// NewSupplierTransform ...
func NewSupplierTransform() SupplierTransformInterface {
	return SupplierTransformSingleton
}
