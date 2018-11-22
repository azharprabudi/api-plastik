package presentations

import (
	"encoding/json"
	"net/http"

	"github.com/api-plastik/errors"

	"github.com/api-plastik/db"
	"github.com/api-plastik/dto"
	jsonParse "github.com/api-plastik/helpers/json"
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

	var validations = []string{}
	itemCatIncReq := new(dto.ItemCategoryIncReq)

	// parse json
	jsonParse.JSONDecode(r.Body, itemCatIncReq)

	// do validations
	if itemCatIncReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		respErr, _ := json.Marshal(errors.NewErrorReponse(errors.ValidationError, "Validation is required", "", validations))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respErr)
		return
	}

	err := item.itemService.CreateItemCategory(itemCatIncReq)
	if err != nil {
		// response error
		respErr, _ := json.Marshal(errors.NewErrorReponse(errors.InternalServerError, err.Error(), "", nil))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(nil)
	return
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
