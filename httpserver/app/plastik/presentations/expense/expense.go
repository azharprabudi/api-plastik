package presentations

import (
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"

	"github.com/azharprabudi/api-plastik/helper/baseurl"
	newError "github.com/azharprabudi/api-plastik/httpserver/error"
	"github.com/azharprabudi/api-plastik/httpserver/request"
	"github.com/azharprabudi/api-plastik/httpserver/response"

	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/internal/expense/dto"
	"github.com/azharprabudi/api-plastik/internal/expense/service"
)

// Find ...
func (e *Expense) Find(w http.ResponseWriter, r *http.Request) {
	results, err := e.service.GetExpenses()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, results)
	return
}

// FindByID ...
func (e *Expense) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result, err := e.service.GetExpenseByID(id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (e *Expense) Create(w http.ResponseWriter, r *http.Request) {
	var validations = []string{}
	req := new(dto.ExpenseReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.ExpenseTypeID == uuid.Nil {
		validations = append(validations, "expense type id field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := e.service.CreateExpense(req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "expense", id),
	}
	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// NewExpensePresentation ...
func NewExpensePresentation(db *db.DB) ExpenseAbstract {
	return &Expense{
		service: service.NewExpenseService(db),
	}
}
