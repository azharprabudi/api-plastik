package presentations

import (
	"net/http"

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
func (item *Item) Find(w http.ResponseWriter, r *http.Request) {
	results, err := item.itemService.GetItem()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (item *Item) FindByID(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "id")

	result := item.itemService.GetItemByID(itemID)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (item *Item) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	itemReq := new(dto.ItemReq)

	// parse json
	request.Get(r.Body, itemReq)

	// do validations
	if itemReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := item.itemService.CreateItem(itemReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "item/", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (item *Item) Update(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	itemID := chi.URLParam(r, "id")

	var validations = []string{}
	itemReq := new(dto.ItemReq)

	// parse json
	request.Get(r.Body, itemReq)

	// do validations
	if itemReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err := item.itemService.UpdateItem(itemID, itemReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (item *Item) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	itemID := chi.URLParam(r, "id")

	err := item.itemService.DeleteItem(itemID)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewItemPresentation ...
func NewItemPresentation(db *db.DB) presentations.BaseAbstract {
	return &Item{
		itemService: service.NewItemService(db),
	}
}
