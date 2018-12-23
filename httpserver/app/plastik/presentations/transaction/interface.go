package presentations

import "net/http"

// TransactionInterface ...
type TransactionInterface interface {
	Find(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	CreateTransactionIn(w http.ResponseWriter, r *http.Request)
	CreateTransactionOut(w http.ResponseWriter, r *http.Request)
	CreateTransactionEtc(w http.ResponseWriter, r *http.Request)
	CreateTransactionEtcType(w http.ResponseWriter, r *http.Request)
	FindTransactionEtcTypes(w http.ResponseWriter, r *http.Request)
	FindTransactionEtcTypeByID(w http.ResponseWriter, r *http.Request)
	UpdateTransactionEtcType(w http.ResponseWriter, r *http.Request)
	DeleteTransactionEtcType(w http.ResponseWriter, r *http.Request)
}
