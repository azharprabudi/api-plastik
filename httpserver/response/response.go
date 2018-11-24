package response

import (
	"encoding/json"
	"io"
)

// JSONEncode ...
func JSONEncode(w io.WriteCloser, model interface{}) error {
	err := json.NewEncoder(w).Encode(model)
	return err
}
