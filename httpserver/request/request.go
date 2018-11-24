package request

import (
	"encoding/json"
	"io"
)

// GetRequest ...
func GetRequest(r io.Reader, model interface{}) error {
	err := json.NewDecoder(r).Decode(model)
	return err
}
