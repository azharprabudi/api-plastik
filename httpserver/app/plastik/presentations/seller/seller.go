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
	results, err := s.service.GetSellers()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, results)
	return
}

// FindByID ...
func (s *Seller) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result, err := s.service.GetSellerByID(id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (s *Seller) Create(w http.ResponseWriter, r *http.Request) {
	req := new(dto.SellerReq)
	var validations = []string{}
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := s.service.CreateSeller(req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers and return
	headers := map[string]string{
		"location": baseurl.Get(r, "sellers", id),
	}
	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (s *Seller) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	req := new(dto.SellerReq)
	var validations = []string{}
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = s.service.UpdateSeller(id, req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (s *Seller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = s.service.DeleteSeller(id)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewSellerPresentation ...
func NewSellerPresentation(db *db.DB) presentations.BaseInterface {
	return &Seller{
		service: service.NewSellerService(db),
	}
}
