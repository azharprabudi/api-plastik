package middlewares

import (
	"encoding/json"
	"net/http"

	newError "github.com/azharprabudi/api-plastik/httpserver/error"
)

// CheckContentType ...
func CheckContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Content-Type"]) > 0 && r.Header["Content-Type"][0] == "application/json" {
			next.ServeHTTP(w, r)
			return
		}

		// create error response
		err, _ := json.Marshal(newError.NewErrorReponse(newError.InternalServerError, "Need content type", "", nil))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Accept", "application/json")
		w.Write(err)
	})
}
