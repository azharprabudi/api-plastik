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
	"github.com/api-plastik/internal/expense/dto"
	"github.com/api-plastik/internal/expense/service"
)

// Find ...
func (expense *Expense) Find(w http.ResponseWriter, r *http.Request) {
	results, err := expense.service.GetExpense()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (expense *Expense) FindByID(w http.ResponseWriter, r *http.Request) {
	expenseID := chi.URLParam(r, "id")

	result := expense.service.GetExpenseByID(expenseID)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (expense *Expense) Create(w http.ResponseWriter, r *http.Request) {

	var validations = []string{}
	expenseReq := new(dto.ExpenseReq)

	// parse json
	request.Get(r.Body, expenseReq)

	// do validations
	if expenseReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	if expenseReq.Amount == "" {
		validations = append(validations, "amount field is required")
	}

	if expenseReq.ExpenseTypeID == "" {
		validations = append(validations, "expenseTypeId field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := expense.service.CreateExpense(expenseReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "expense/", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// Update ...
func (expense *Expense) Update(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	expenseID := chi.URLParam(r, "id")

	var validations = []string{}
	expenseReq := new(dto.ExpenseReq)

	// parse json
	request.Get(r.Body, expenseReq)

	// do validations
	if expenseReq.Name == "" {
		validations = append(validations, "name field is required")
	}

	if expenseReq.Amount == "" {
		validations = append(validations, "amount field is required")
	}

	if expenseReq.ExpenseTypeID == "" {
		validations = append(validations, "expense_type_id field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err := expense.service.UpdateExpense(expenseID, expenseReq)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// Delete ...
func (expense *Expense) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from url parameter
	expenseID := chi.URLParam(r, "id")

	err := expense.service.DeleteExpense(expenseID)
	if err != nil {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewExpensePresentation ...
func NewExpensePresentation(db *db.DB) presentations.BaseAbstract {
	return &Expense{
		service: service.NewExpenseService(db),
	}
}
