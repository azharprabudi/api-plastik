package presentations

import (
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/go-chi/chi"

	"github.com/api-plastik/helper/baseurl"
	"github.com/api-plastik/httpserver/app/plastik/presentations"
	newError "github.com/api-plastik/httpserver/error"
	"github.com/api-plastik/httpserver/request"
	"github.com/api-plastik/httpserver/response"

	"github.com/api-plastik/db"
	"github.com/api-plastik/internal/seller/dto"
	"github.com/api-plastik/internal/seller/service"
)

// Find ...
func (seller *Seller) Find(w http.ResponseWriter, r *http.Request) {
	results, err := seller.sellerService.GetSeller()
	if err != nil {
		response.SendResponse(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.SendResponse(w, http.StatusOK, nil, results)
}

// FindByID ...
func (seller *Seller) FindByID(w http.ResponseWriter, r *http.Request) {
	// get query param id
	sellerID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(sellerID)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := seller.sellerService.GetSellerByID(u)
	response.SendResponse(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (seller *Seller) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	sellerCatIncReq := new(dto.SellerReq)

	// parse json
	request.GetRequest(r.Body, sellerCatIncReq)

	// do validations
	if sellerCatIncReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	if sellerCatIncReq.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.SendResponse(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := seller.sellerService.CreateSeller(sellerCatIncReq)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "sellercategories/", id),
	}

	response.SendResponse(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (seller *Seller) Update(w http.ResponseWriter, r *http.Request) {
	// get query param id
	sellerID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(sellerID)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	sellerReq := new(dto.SellerReq)

	// parse json
	request.GetRequest(r.Body, sellerReq)

	// do validations
	if sellerReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.SendResponse(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = seller.sellerService.UpdateSeller(u, sellerReq)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.SendResponse(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (seller *Seller) Delete(w http.ResponseWriter, r *http.Request) {
	// get query param id
	sellerID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(sellerID)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = seller.sellerService.DeleteSeller(u)
	if err != nil {
		// response error
		response.SendResponse(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.SendResponse(w, http.StatusOK, nil, nil)
	return
}

// NewPresentationSeller ...
func NewPresentationSeller(db *db.DB) presentations.BaseAbstract {
	return &Seller{
		sellerService: service.NewSellerService(db),
	}
}
