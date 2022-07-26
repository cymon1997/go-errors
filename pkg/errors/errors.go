package errors

import "net/http"

type Error struct {
	message string
	code    int
}

// New create new error with predefined error code
// recommended to set code as http status code
func New(code int) *Error {
	return &Error{
		code: code,
	}
}

// WithMessage specifies custom error message
func (e *Error) WithMessage(message string) *Error {
	e.message = message
	return e
}

// Code return the error code
// recommended to set code as http status code
func (e *Error) Code() int {
	return e.code
}

// Error return the error message
func (e *Error) Error() string {
	if e.message != "" {
		return e.message
	}
	return http.StatusText(e.code)
}

// IsShouldRetry decide whether the error should be retried or not
// retryable error is from server error (5xx)
func (e *Error) IsShouldRetry() bool {
	return e.code >= 500
}
