package transform

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/azharprabudi/api-plastik/internal/supplier/model"
)

// MakeModelCreateSupplier ...
func (st *SupplierTransform) MakeModelCreateSupplier(companyID uuid.UUID, req *dto.SupplierReq) *model.SupplierCreate {
	return &model.SupplierCreate{
		Supplier: model.Supplier{
			SupplierID: uuid.NewV4(),
			Name:       req.Supplier.Name,
			Address:    req.Supplier.Address,
			Phone:      req.Supplier.Phone,
			CreatedAt:  time.Now().UTC(),
			CompanyID:  companyID,
		},
	}
}

// MakeModelUpdateSupplier ...
func (st *SupplierTransform) MakeModelUpdateSupplier(req *dto.SupplierReq) *model.SupplierUpdate {
	return &model.SupplierUpdate{
		Name:    req.Supplier.Name,
		Address: req.Supplier.Address,
		Phone:   req.Supplier.Phone,
	}
}

// MakeResponseGetSuppliers ...
func (st *SupplierTransform) MakeResponseGetSuppliers(res []*model.SupplierRead) []*dto.SupplierRes {
	var results = []*dto.SupplierRes{}
	for _, supplier := range res {
		results = append(results, &dto.SupplierRes{
			ID:        supplier.Supplier.SupplierID,
			CreatedAt: supplier.Supplier.CreatedAt,
			Supplier: dto.Supplier{
				Name:    supplier.Supplier.Name,
				Phone:   supplier.Supplier.Phone,
				Address: supplier.Supplier.Address,
			},
		})
	}

	return results
}

// MakeResponseGetSupplierByID ...
func (st *SupplierTransform) MakeResponseGetSupplierByID(res *model.SupplierRead) *dto.SupplierRes {
	return &dto.SupplierRes{
		ID:        res.SupplierID,
		CreatedAt: res.CreatedAt,
		Supplier: dto.Supplier{
			Name:    res.Name,
			Phone:   res.Phone,
			Address: res.Address,
		},
	}
}

// NewSupplierTransform ...
func NewSupplierTransform() SupplierTransformInterface {
	return SupplierTransformSingleton
}
