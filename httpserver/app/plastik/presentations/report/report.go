package presentations

import (
	"net/http"

	"github.com/azharprabudi/api-plastik/db"
	newError "github.com/azharprabudi/api-plastik/httpserver/error"
	"github.com/azharprabudi/api-plastik/httpserver/response"
	"github.com/azharprabudi/api-plastik/internal/transaction/service"

	itemService "github.com/azharprabudi/api-plastik/internal/item/service"
	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
)

// GetCountTransactions ...
func (report *Report) GetCountTransactions(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	endAt := r.URL.Query().Get("endAt")
	startAt := r.URL.Query().Get("startAt")

	count, err := report.transactionService.GetCountTransactions(companyID, startAt, endAt)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, struct {
		Value int `json:"value"`
	}{
		Value: count,
	})
	return
}

// GetSummaryTransactions ...
func (report *Report) GetSummaryTransactions(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	endAt := r.URL.Query().Get("endAt")
	startAt := r.URL.Query().Get("startAt")

	amount, err := report.transactionService.GetSummaryTransactions(companyID, startAt, endAt)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, struct {
		Value float64 `json:"value"`
	}{
		Value: amount,
	})
	return

}

// GetSummaryTransactionsIn ...
func (report *Report) GetSummaryTransactionsIn(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	endAt := r.URL.Query().Get("endAt")
	startAt := r.URL.Query().Get("startAt")

	amount, err := report.transactionService.GetSummaryTransactionsIn(companyID, startAt, endAt)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, struct {
		Value float64 `json:"value"`
	}{
		Value: amount,
	})
}

// GetSummaryTransactionsOut ...
func (report *Report) GetSummaryTransactionsOut(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	endAt := r.URL.Query().Get("endAt")
	startAt := r.URL.Query().Get("startAt")

	amount, err := report.transactionService.GetSummaryTransactionsOut(companyID, startAt, endAt)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, struct {
		Value float64 `json:"value"`
	}{
		Value: amount,
	})
}

// GetSummaryTransactionsEtc ...
func (report *Report) GetSummaryTransactionsEtc(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	endAt := r.URL.Query().Get("endAt")
	startAt := r.URL.Query().Get("startAt")

	amount, err := report.transactionService.GetSummaryTransactionsEtc(companyID, startAt, endAt)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, struct {
		Value float64 `json:"value"`
	}{
		Value: amount,
	})
}

// GetItemStockLogs ...
func (report *Report) GetItemStockLogs(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.FromString(chi.URLParam(r, "companyId"))
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	res, err := report.itemService.GetItemStockLogs(companyID)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, nil, newError.NewErrorReponse(newError.InternalServerError, err.Error(), "", nil))
		return
	}

	response.Send(w, http.StatusOK, nil, res)
}

// NewReportPresentation ...
func NewReportPresentation(db *db.DB) ReportInterface {
	return &Report{
		itemService:        itemService.NewItemService(db),
		transactionService: service.NewTransactionService(db),
	}
}
