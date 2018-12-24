package presentations

import (
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"

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
func (i *Item) Find(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	results, err := i.service.GetItems(companyID)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (i *Item) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result, err := i.service.GetItemByID(companyID, id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (i *Item) Create(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	req := new(dto.ItemReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.ItemCategoryID == uuid.Nil {
		validations = append(validations, "item category id field is required")
	}

	if req.UnitID == uuid.Nil {
		validations = append(validations, "unit id field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := i.service.CreateItem(companyID, req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers and return it
	headers := map[string]string{
		"location": baseurl.Get(r, "item", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (i *Item) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	req := new(dto.ItemReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.ItemCategoryID == uuid.Nil {
		validations = append(validations, "itemCategoryId field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = i.service.UpdateItem(companyID, id, req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (i *Item) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = i.service.DeleteItem(companyID, id)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewItemPresentation ...
func NewItemPresentation(db *db.DB) presentations.BaseInterface {
	return &Item{
		service: service.NewItemService(db),
	}
}
