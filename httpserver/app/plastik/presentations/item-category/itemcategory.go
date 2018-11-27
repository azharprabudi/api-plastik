package presentations

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

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
	results, err := item.itemService.GetItemCategory()
	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.SendResponse(w, http.StatusCreated, results)
}

// FindByID ...
func (item *ItemCategory) FindByID(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result, err := item.itemService.GetItemCategoryByID(categoryID)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.SendResponse(w, http.StatusCreated, result)
	return
}

// Create ...
func (item *ItemCategory) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	itemCatIncReq := new(dto.ItemCategoryIncReq)

	// parse json
	request.GetRequest(r.Body, itemCatIncReq)

	// do validations
	if itemCatIncReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.SendResponse(w, http.StatusBadRequest, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err := item.itemService.CreateItemCategory(itemCatIncReq)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusBadRequest, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.SendResponse(w, http.StatusCreated, nil)
	return
}

// Update ...
func (item *ItemCategory) Update(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	itemCatIncReq := new(dto.ItemCategoryIncReq)

	// parse json
	request.GetRequest(r.Body, itemCatIncReq)

	// do validations
	if itemCatIncReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.SendResponse(w, http.StatusBadRequest, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = item.itemService.UpdateItemCategory(categoryID, itemCatIncReq)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusBadRequest, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.SendResponse(w, http.StatusOK, nil)
	return
}

// Delete ...
func (item *ItemCategory) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	categoryID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = item.itemService.DeleteItemCategory(categoryID)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusBadRequest, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.SendResponse(w, http.StatusOK, nil)
	return
}

// NewPresentationItemCategory ...
func NewPresentationItemCategory(db *db.DB) presentations.BaseAbstract {
	return &ItemCategory{
		itemService: service.NewItemService(db),
	}
}
