package presentations

import (
	"net/http"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/seller/service"
	"github.com/api-plastik/presentations"
)

// Find ...
func (item *Seller) Find(w http.ResponseWriter, r *http.Request) {

}

// FindByID ...
func (item *Seller) FindByID(w http.ResponseWriter, r *http.Request) {

}

// Create ...
func (item *Seller) Create(w http.ResponseWriter, r *http.Request) {

}

// Update ...
func (item *Seller) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ...
func (item *Seller) Delete(w http.ResponseWriter, r *http.Request) {

}

// NewPresentationSeller ...
func NewPresentationSeller(db *db.DB) presentations.BaseAbstract {
	return &Seller{
		sellerService: service.NewSellerService(db),
	}
}
