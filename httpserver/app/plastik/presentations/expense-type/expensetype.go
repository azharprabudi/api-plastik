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
	"github.com/api-plastik/internal/expense/dto"
	"github.com/api-plastik/internal/expense/service"
)

// Find ...
func (expense *ExpenseType) Find(w http.ResponseWriter, r *http.Request) {
	results, err := expense.service.GetExpenseType()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (expense *ExpenseType) FindByID(w http.ResponseWriter, r *http.Request) {
	expenseTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := expense.service.GetExpenseTypeByID(expenseTypeID)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (expense *ExpenseType) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	expenseTypeReq := new(dto.ExpenseTypeReq)

	// parse json
	request.Get(r.Body, expenseTypeReq)

	// do validations
	if expenseTypeReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := expense.service.CreateExpenseType(expenseTypeReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "expensetype/", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (expense *ExpenseType) Update(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	expenseTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	expenseTypeReq := new(dto.ExpenseTypeReq)

	// parse json
	request.Get(r.Body, expenseTypeReq)

	// do validations
	if expenseTypeReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = expense.service.UpdateExpenseType(expenseTypeID, expenseTypeReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (expense *ExpenseType) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	expenseTypeID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		// response error
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = expense.service.DeleteExpenseType(expenseTypeID)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewPresentationExpenseType ...
func NewPresentationExpenseType(db *db.DB) presentations.BaseAbstract {
	return &ExpenseType{
		service: service.NewExpenseTypeService(db),
	}
}
