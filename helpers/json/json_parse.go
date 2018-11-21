package helpers

import (
	"encoding/json"
	"io"
)

// JSONDecode ...
func JSONDecode(r io.Reader, model interface{}) error {
	err := json.NewDecoder(r).Decode(model)
	return err
}

// JSONEncode ...
func JSONEncode(w io.WriteCloser, model interface{}) error {
	err := json.NewEncoder(w).Encode(model)
	return err
}
