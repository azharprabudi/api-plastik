package presentations

import (
	"encoding/json"
	"net/http"

	"github.com/api-plastik/httpserver/app/plastik/presentations"
	newError "github.com/api-plastik/httpserver/error"
	"github.com/api-plastik/httpserver/request"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/service"
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
	request.JSONDecode(r.Body, itemCatIncReq)

	// do validations
	if itemCatIncReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		respErr, _ := json.Marshal(newError.NewErrorReponse(newError.ValidationError, "Validation is required", "", validations))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respErr)
		return
	}

	err := item.itemService.CreateItemCategory(itemCatIncReq)
	if err != nil {
		// response error
		respErr, _ := json.Marshal(newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
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
