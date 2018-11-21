package presentations

import (
	"net/http"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/service"
	"github.com/api-plastik/presentations"
)

// Find ...
func (item *ItemCategory) Find(w http.ResponseWriter, r *http.Request) {

}

// FindByID ...
func (item *ItemCategory) FindByID(w http.ResponseWriter, r *http.Request) {

}

// Create ...
func (item *ItemCategory) Create(w http.ResponseWriter, r *http.Request) {

}

// Update ...
func (item *ItemCategory) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ...
func (item *ItemCategory) Delete(w http.ResponseWriter, r *http.Request) {

}

// NewPresentationItemCategory ...
func NewPresentationItemCategory(db *db.DB) presentations.BaseAbstract {
	return &ItemCategory{
		itemService: service.NewItemService(db),
	}
}
