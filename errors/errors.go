package errors

import "net/http"

type Error struct {
	code    string
	message string
	status  int
}

// New create new error with predefined status
// It's recommended to set status as http status code
func New(status int) *Error {
	return &Error{
		status: status,
	}
}

// WithCode specifies custom error code
func (e *Error) WithCode(code string) *Error {
	e.code = code
	return e
}

// WithMessage specifies custom error message
func (e *Error) WithMessage(msg string) *Error {
	e.message = msg
	return e
}

// Code return the error code
// recommended to set code as http status code
func (e *Error) Code() string {
	return e.code
}

// Message return the error code
// recommended to set code as http status code
func (e *Error) Message() string {
	return e.message
}

// Status return the error code
// recommended to set code as http status code
func (e *Error) Status() int {
	return e.status
}

// Error return the error message
func (e *Error) Error() string {
	if e.message != "" {
		return e.message
	}
	return http.StatusText(e.status)
}

// IsRetry decide whether the error should be retried or not
// retryable error is from server error (5xx)
func (e *Error) IsRetry() bool {
	return e.status >= 500 && e.status < 600
}
