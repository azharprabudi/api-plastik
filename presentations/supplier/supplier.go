package presentations

import (
	"net/http"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/supplier/service"
	"github.com/api-plastik/presentations"
)

// Find ...
func (item *Supplier) Find(w http.ResponseWriter, r *http.Request) {

}

// FindByID ...
func (item *Supplier) FindByID(w http.ResponseWriter, r *http.Request) {

}

// Create ...
func (item *Supplier) Create(w http.ResponseWriter, r *http.Request) {

}

// Update ...
func (item *Supplier) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ...
func (item *Supplier) Delete(w http.ResponseWriter, r *http.Request) {

}

// NewPresentationSupplier ...
func NewPresentationSupplier(db *db.DB) presentations.BaseAbstract {
	return &Supplier{
		supplierService: service.NewSupplierService(db),
	}
}
