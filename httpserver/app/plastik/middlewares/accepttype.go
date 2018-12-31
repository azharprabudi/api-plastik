package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"

	newError "github.com/azharprabudi/api-plastik/httpserver/error"
)

// AcceptContentType ...
func AcceptContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Accept"]) > 0 && strings.Contains(r.Header["Accept"][0], "application/json") {
			next.ServeHTTP(w, r)
			return
		}

		// create error response
		err, _ := json.Marshal(newError.NewErrorReponse(newError.InternalServerError, "We just accept the content type json", "", nil))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(err)
	})
}
