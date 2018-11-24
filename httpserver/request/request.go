package request

import (
	"encoding/json"
	"io"
)

// JSONDecode ...
func JSONDecode(r io.Reader, model interface{}) error {
	err := json.NewDecoder(r).Decode(model)
	return err
}
