package presentations

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/api-plastik/helper/baseurl"
	"github.com/api-plastik/httpserver/app/plastik/presentations"
	newError "github.com/api-plastik/httpserver/error"
	"github.com/api-plastik/httpserver/request"
	"github.com/api-plastik/httpserver/response"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/item/dto"
	"github.com/api-plastik/internal/item/service"
)

// Find ...
func (item *ItemCategory) Find(w http.ResponseWriter, r *http.Request) {
	results, err := item.service.GetItemCategory()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (item *ItemCategory) FindByID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := item.service.GetItemCategoryByID(categoryID)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (item *ItemCategory) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	itemCatReq := new(dto.ItemCategoryReq)

	// parse json
	request.Get(r.Body, itemCatReq)

	// do validations
	if itemCatReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := item.service.CreateItemCategory(itemCatReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "itemcategories/", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (item *ItemCategory) Update(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	itemCatReq := new(dto.ItemCategoryReq)

	// parse json
	request.Get(r.Body, itemCatReq)

	// do validations
	if itemCatReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = item.service.UpdateItemCategory(categoryID, itemCatReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (item *ItemCategory) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = item.service.DeleteItemCategory(categoryID)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewPresentationItemCategory ...
func NewPresentationItemCategory(db *db.DB) presentations.BaseAbstract {
	return &ItemCategory{
		service: service.NewItemService(db),
	}
}
