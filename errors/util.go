package errors

import (
	"net/http"
)

func Is(err, target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return err.Error() == target.Error()
	}
	e, ok := err.(*Error)
	if !ok {
		return err.Error() == t.Error()
	}
	return e.Code() == t.Code() &&
		e.Error() == t.Error()
}

func GetCode(err error) string {
	errs, ok := err.(*Error)
	if ok {
		return errs.Code()
	}
	return "UNKNOWN"
}

func GetMessage(err error) string {
	errs, ok := err.(*Error)
	if ok {
		return errs.Message()
	}
	return err.Error()
}

func GetStatus(err error) int {
	errs, ok := err.(*Error)
	if ok {
		return errs.Status()
	}
	return http.StatusInternalServerError
}

// IsRetry check if error must be retried, default true
// default is false to reduce unnecessary retry
func IsRetry(err error) bool {
	errs, ok := err.(*Error)
	if ok {
		return errs.IsRetry()
	}
	return false
}
