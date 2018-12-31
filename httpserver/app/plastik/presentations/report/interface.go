package presentations

import "net/http"

// ReportInterface ...
type ReportInterface interface {
	GetCountTransactions(w http.ResponseWriter, r *http.Request)
	GetSummaryTransactions(w http.ResponseWriter, r *http.Request)
	GetSummaryTransactionsIn(w http.ResponseWriter, r *http.Request)
	GetSummaryTransactionsOut(w http.ResponseWriter, r *http.Request)
	GetSummaryTransactionsEtc(w http.ResponseWriter, r *http.Request)
	GetItemStockLogs(w http.ResponseWriter, r *http.Request)
}
