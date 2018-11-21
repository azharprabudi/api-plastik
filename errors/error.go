package errors

// NewError ...
func NewError(code string, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
