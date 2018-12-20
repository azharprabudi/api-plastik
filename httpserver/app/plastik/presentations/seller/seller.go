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
	"github.com/azharprabudi/api-plastik/internal/seller/dto"
	"github.com/azharprabudi/api-plastik/internal/seller/service"
)

// Find ...
func (s *Seller) Find(w http.ResponseWriter, r *http.Request) {
	results, err := s.service.GetSeller()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (s *Seller) FindByID(w http.ResponseWriter, r *http.Request) {
	// get query param id
	sellerID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(sellerID)
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := s.service.GetSellerByID(u)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (s *Seller) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	seller := new(dto.SellerReq)

	// parse json
	request.Get(r.Body, seller)

	// do validations
	if seller.Name == "" {
		validations = append(validations, "name field is required")
	}

	if seller.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := s.service.CreateSeller(seller)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "sellers", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (s *Seller) Update(w http.ResponseWriter, r *http.Request) {
	// get query param id
	sellerID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(sellerID)
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	seller := new(dto.SellerReq)

	// parse json
	request.Get(r.Body, seller)

	// do validations
	if seller.Name == "" {
		validations = append(validations, "name field is required")
	}

	if seller.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = s.service.UpdateSeller(u, seller)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (s *Seller) Delete(w http.ResponseWriter, r *http.Request) {
	// get query param id
	sellerID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(sellerID)
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = s.service.DeleteSeller(u)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewPresentationSeller ...
func NewPresentationSeller(db *db.DB) presentations.BaseInterface {
	return &Seller{
		service: service.NewSellerService(db),
	}
}
