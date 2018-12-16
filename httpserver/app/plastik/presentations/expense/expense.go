package presentations

import (
	"net/http"

	"github.com/satori/go.uuid"

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
func (e *Expense) Find(w http.ResponseWriter, r *http.Request) {
	results, err := expense.service.GetExpense()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}
	response.Send(w, http.StatusOK, nil, results)
}

// FindByID ...
func (e *Expense) FindByID(w http.ResponseWriter, r *http.Request) {
	expenseID, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result := expense.service.GetExpenseByID(expenseID)
	response.Send(w, http.StatusOK, nil, result)
	return
}

// Create ...
func (e *Expense) Create(w http.ResponseWriter, r *http.Request) {
	var validations = []string{}
	req := new(dto.ExpenseReq)

	// parse json
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if req.Amount == "" {
		validations = append(validations, "amount field is required")
	}

	if req.ExpenseTypeID == uuid.Nil {
		validations = append(validations, "expense type id field is required")
	}

	// if validation exists there is error
	if len(validations) > 0 {
		// response error
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := expense.service.CreateExpense(req)
	if err != nil {
		// response error
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
func NewExpensePresentation(db *db.DB) presentations.ExpenseAbstract {
	return &Expense{
		service: service.NewExpenseService(db),
	}
}
