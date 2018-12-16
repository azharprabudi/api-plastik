package presentations

import "net/http"

// ExpenseAbstract ...
type ExpenseAbstract interface {
	Find(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}
