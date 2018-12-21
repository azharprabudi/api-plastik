package presentations

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"

	"github.com/azharprabudi/api-plastik/helper/baseurl"
	"github.com/azharprabudi/api-plastik/httpserver/request"
	"github.com/azharprabudi/api-plastik/httpserver/response"

	"github.com/azharprabudi/api-plastik/db"
	newError "github.com/azharprabudi/api-plastik/httpserver/error"
	itemService "github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/azharprabudi/api-plastik/internal/transaction/dto"
	transactionService "github.com/azharprabudi/api-plastik/internal/transaction/service"
	"github.com/azharprabudi/api-plastik/internal/transaction/value"
)

// Find ...
func (t *Transaction) Find(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	endAt := r.URL.Query().Get("endAt")
	startAt := r.URL.Query().Get("startAt")
	orderBy := r.URL.Query().Get("orderBy")

	transactions, err := t.service.FindTransactions(page, startAt, endAt, orderBy)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusCreated, nil, transactions)
	return
}

// Create ...
func (t *Transaction) Create(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionReq)
	request.Get(r.Body, req)

	// do some validation
	var validations []string
	if req.Type != value.TRANSACTION_IN && req.Type != value.TRANSACTION_OUT {
		validations = append(validations, "type is not match")
	}

	if req.SupplierID == nil && req.SellerID == nil {
		validations = append(validations, "seller id or supplier id is required")
	}

	if len(req.Details) < 1 {
		validations = append(validations, "details must more than one")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	// create a new transaction
	id, err := t.service.CreateTransaction(req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create header and return it
	headers := map[string]string{
		"location": baseurl.Get(r, "transaction", id),
	}

	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// FindByID ...
func (t *Transaction) FindByID(w http.ResponseWriter, r *http.Request) {
	id := uuid.FromString(chi.URLParam(r, id))
	transaction, err := t.service.FindTransactionByID(id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusCreated, nil, transaction)
	return
}

// NewPresentationSupplier ...
func NewPresentationTransaction(db *db.DB) TransactionInterface {
	itemService := itemService.NewItemService(db)
	return &Transaction{
		service: transactionService.NewTransactionService(db, itemService),
	}
}
