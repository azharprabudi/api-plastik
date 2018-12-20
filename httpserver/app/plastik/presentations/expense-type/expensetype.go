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
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/service"
)

// Find ...
func (et *ExpenseType) Find(w http.ResponseWriter, r *http.Request) {
	results, err := et.service.GetExpenseType()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (et *ExpenseType) FindByID(w http.ResponseWriter, r *http.Request) {
	expenseTypeID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := et.service.GetExpenseTypeByID(expenseTypeID)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (et *ExpenseType) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	req := new(dto.ExpenseTypeReq)

	// parse json
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := et.service.CreateExpenseType(req)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "expense-type", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (et *ExpenseType) Update(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	expenseTypeID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	req := new(dto.ExpenseTypeReq)

	// parse json
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = et.service.UpdateExpenseType(expenseTypeID, req)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (et *ExpenseType) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	expenseTypeID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = et.service.DeleteExpenseType(expenseTypeID)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewExpenseTypePresentation ...
func NewExpenseTypePresentation(db *db.DB) presentations.BaseInterface {
	return &ExpenseType{
		service: service.NewExpenseService(db),
	}
}
