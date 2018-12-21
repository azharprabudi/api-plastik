package presentations

import (
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/helper/baseurl"
	"github.com/azharprabudi/api-plastik/httpserver/app/plastik/presentations"
	newError "github.com/azharprabudi/api-plastik/httpserver/error"
	"github.com/azharprabudi/api-plastik/httpserver/request"
	"github.com/azharprabudi/api-plastik/httpserver/response"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/item/dto"
	"github.com/azharprabudi/api-plastik/internal/item/service"
)

// Find ...
func (ic *ItemCategory) Find(w http.ResponseWriter, r *http.Request) {
	results, err := ic.service.GetItemCategories()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (ic *ItemCategory) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result, err := ic.service.GetItemCategoryByID(id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (ic *ItemCategory) Create(w http.ResponseWriter, r *http.Request) {
	var validations = []string{}
	req := new(dto.ItemCategoryReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := ic.service.CreateItemCategory(req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers and return it
	headers := map[string]string{
		"location": baseurl.Get(r, "item-category", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (ic *ItemCategory) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	req := new(dto.ItemCategoryReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = ic.service.UpdateItemCategory(id, req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (ic *ItemCategory) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = ic.service.DeleteItemCategory(id)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewPresentationItemCategory ...
func NewPresentationItemCategory(db *db.DB) presentations.BaseInterface {
	return &ItemCategory{
		service: service.NewItemService(db),
	}
}
