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

	response.Send(w, http.StatusOK, nil, transactions)
	return
}

// CreateTransactionIn ...
func (t *Transaction) CreateTransactionIn(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionInReq)
	request.Get(r.Body, req)

	// do some validation
	var validations []string
	if req.SupplierID == uuid.Nil {
		validations = append(validations, "supplier id is required")
	}

	if len(req.Details) < 1 {
		validations = append(validations, "details must more than one")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	// create a new transaction
	id, err := t.service.CreateTransactionIn(req, value.TRANSACTION_IN)
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

// CreateTransactionOut ...
func (t *Transaction) CreateTransactionOut(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionOutReq)
	request.Get(r.Body, req)

	// do some validation
	var validations []string
	if req.SellerID == uuid.Nil {
		validations = append(validations, "seller id is required")
	}

	if len(req.Details) < 1 {
		validations = append(validations, "details must more than one")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	// create a new transaction
	id, err := t.service.CreateTransactionOut(req, value.TRANSACTION_OUT)
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

// CreateTransactionEtc ...
func (t *Transaction) CreateTransactionEtc(w http.ResponseWriter, r *http.Request) {
	req := new(dto.TransactionEtcReq)
	request.Get(r.Body, req)

	// do some validation
	var validations []string
	if req.TransactionEtc.Amount == 0 {
		validations = append(validations, "amount more than zero")
	}

	if req.TransactionEtc.TransactionEtcTypeID == uuid.Nil {
		validations = append(validations, "transaction type is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	// create a new transaction
	id, err := t.service.CreateTransactionEtc(req, value.TRANSACTION_ETC)
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
	id, _ := uuid.FromString(chi.URLParam(r, "id"))
	transaction, err := t.service.FindTransactionByID(id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusCreated, nil, transaction)
	return
}

// FindTransactionEtcTypes ...
func (t *Transaction) FindTransactionEtcTypes(w http.ResponseWriter, r *http.Request) {
	results, err := t.service.FindTransactionEtcTypes()
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, results)
	return
}

// FindTransactionEtcTypeByID ...
func (t *Transaction) FindTransactionEtcTypeByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	result, err := t.service.FindTransactionEtcTypeByID(id)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, result)
	return
}

// CreateTransactionEtcType ...
func (t *Transaction) CreateTransactionEtcType(w http.ResponseWriter, r *http.Request) {
	var validations = []string{}
	req := new(dto.TransactionEtcTypeReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	id, err := t.service.CreateTransactionEtcType(req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	// create headers
	headers := map[string]string{
		"location": baseurl.Get(r, "transaction/etc/type", id),
	}
	response.Send(w, http.StatusCreated, headers, nil)
	return
}

// UpdateTransactionEtcType ...
func (t *Transaction) UpdateTransactionEtcType(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	var validations = []string{}
	req := new(dto.TransactionEtcTypeReq)
	request.Get(r.Body, req)

	// do validations
	if req.Name == "" {
		validations = append(validations, "name field is required")
	}

	if len(validations) > 0 {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, "", "", validations))
		return
	}

	err = t.service.UpdateTransactionEtcType(id, req)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// DeleteTransactionEtcType ...
func (t *Transaction) DeleteTransactionEtcType(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.FromString(chi.URLParam(r, "id"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	err = t.service.DeleteTransactionEtcType(id)
	if err != nil {
		response.Send(w, http.StatusBadRequest, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, nil)
	return
}

// NewTransactionPresentation ...
func NewTransactionPresentation(db *db.DB) TransactionInterface {
	itemService := itemService.NewItemService(db)
	return &Transaction{
		service: transactionService.NewTransactionService(db, itemService),
	}
}
