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
	"github.com/azharprabudi/api-plastik/internal/supplier/dto"
	"github.com/azharprabudi/api-plastik/internal/supplier/service"
)

// Find ...
func (s *Supplier) Find(w http.ResponseWriter, r *http.Request) {
	results, err := s.service.GetSupplier()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (s *Supplier) FindByID(w http.ResponseWriter, r *http.Request) {
	// get query param id
	supplierID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(supplierID)
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := s.service.GetSupplierByID(u)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (s *Supplier) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	req := new(dto.SupplierReq)

	// parse json
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := s.service.CreateSupplier(req)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "suppliers", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (s *Supplier) Update(w http.ResponseWriter, r *http.Request) {
	// get query param id
	supplierID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(supplierID)
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	supplierReq := new(dto.SupplierReq)

	// parse json
	request.Get(r.Body, supplierReq)

	// do validations
	if supplierReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	if supplierReq.Phone == "" {
		validations = append(validations, "phone field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = s.service.UpdateSupplier(u, supplierReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (s *Supplier) Delete(w http.ResponseWriter, r *http.Request) {
	// get query param id
	supplierID := chi.URLParam(r, "id")

	// parse string to uuid
	u, err := uuid.FromString(supplierID)
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = s.service.DeleteSupplier(u)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewPresentationSupplier ...
func NewPresentationSupplier(db *db.DB) presentations.BaseInterface {
	return &Supplier{
		service: service.NewSupplierService(db),
	}
}
