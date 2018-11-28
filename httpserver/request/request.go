package request

import (
	"encoding/json"
	"io"
)

// Get ...
func Get(r io.Reader, model interface{}) error {
	err := json.NewDecoder(r).Decode(model)
	return err
}
