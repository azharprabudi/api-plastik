package response

import (
	"encoding/json"
	"net/http"
)

// Send ...
func Send(w http.ResponseWriter, statusCode int, headers map[string]string, model interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// check headers is exists or not
	if len(headers) > 0 {
		for headerKey, headerVal := range headers {
			w.Header().Set(headerKey, headerVal)
		}
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model)
}
