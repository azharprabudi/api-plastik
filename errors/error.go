package errors

import (
	"net/http"

	"github.com/go-chi/render"
)

// Render ...
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewError ...
func NewError(code string, message string, statusText string) render.Renderer {
	return &ErrorResponse{
		Code:       code,
		Message:    message,
		StatusText: statusText,
	}
}
