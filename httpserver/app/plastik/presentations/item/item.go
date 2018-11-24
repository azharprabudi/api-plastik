package presentations

import (
	"net/http"

	"github.com/api-plastik/db"
	"github.com/api-plastik/httpserver/app/plastik/presentations"
	itemService "github.com/api-plastik/internal/item/service"
)

// Find ...
func (item *ItemPresentation) Find(w http.ResponseWriter, r *http.Request) {

}

// FindByID ...
func (item *ItemPresentation) FindByID(w http.ResponseWriter, r *http.Request) {

}

// Create ...
func (item *ItemPresentation) Create(w http.ResponseWriter, r *http.Request) {

}

// Update ...
func (item *ItemPresentation) Update(w http.ResponseWriter, r *http.Request) {

}

// Delete ...
func (item *ItemPresentation) Delete(w http.ResponseWriter, r *http.Request) {

}

// NewItemPresentation ...
func NewItemPresentation(db *db.DB) presentations.BaseAbstract {
	return &ItemPresentation{
		itemService: itemService.NewItemService(db),
	}
}
