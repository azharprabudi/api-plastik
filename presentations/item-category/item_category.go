package presentations

import (
	"net/http"

	"github.com/api-plastik/dto"
	"github.com/api-plastik/errors"
	"github.com/api-plastik/helpers/json"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/service"
	"github.com/api-plastik/presentations"
	"github.com/go-chi/render"
)

// Find ...
func (item *ItemCategory) Find(w http.ResponseWriter, r *http.Request) {

}

// FindByID ...
func (item *ItemCategory) FindByID(w http.ResponseWriter, r *http.Request) {

}

// Create ...
func (item *ItemCategory) Create(w http.ResponseWriter, r *http.Request) {
	itemCatIncReq := new(dto.ItemCategoryIncReq)
	err := helpers.JSONDecode(r.Body, itemCatIncReq)

	// validation parsing all
	render.Status(r, http.StatusBadRequest)
	render.Render(w, r, errors.NewError(errors.ValidationError, "Validation is required"))

	// if err != nil {
	// 	return
	// }

	// if itemCatIncReq.Name == "" {
	// 	render.Status(r, http.StatusBadRequest)
	// 	render.Render(w, r, errors.NewError(errors.ValidationError, "Validation is required"))
	// 	return
	// }
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
