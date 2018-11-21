package errors

// NewErrorReponse ...
func NewErrorReponse(code string, message string, statusText string) *ErrorResponse {
	return &ErrorResponse{
		Code:       code,
		Message:    message,
		StatusText: statusText,
	}
}
