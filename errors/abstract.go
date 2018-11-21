package errors

// ErrorResponse ...
type ErrorResponse struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	StatusText string   `json:"statusText"`
	Fields     []string `json:"fields"`
}
