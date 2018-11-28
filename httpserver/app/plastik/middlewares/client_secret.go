package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/api-plastik/config"
	newError "github.com/api-plastik/httpserver/error"
)

// CheckClientSecret ...
func CheckClientSecret(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header["Client_secret"]) > 0 && r.Header["Client_secret"][0] == config.CLIENTSECRET {
			next.ServeHTTP(w, r)
			return
		}

		// create error response
		err, _ := json.Marshal(newError.NewErrorReponse(newError.AuthError, "There is no client secret exists", "", nil))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(err)
	})
}
