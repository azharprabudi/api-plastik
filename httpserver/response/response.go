package response

import (
	"encoding/json"
	"net/http"
)

// SendResponse ...
func SendResponse(w http.ResponseWriter, statusCode int, model interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model)
}
