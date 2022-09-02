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

func GetMessage(err error) string {
	errs, ok := err.(*Error)
	if !ok {
		return err.Error()
	}
	return errs.Error()
}

func GetStatus(err error) int {
	errs, ok := err.(*Error)
	if ok {
		return errs.Code()
	}
	return http.StatusInternalServerError
}

// IsShouldRetry check if error must be retried, default true
// default is false to reduce unnecessary retry
func IsShouldRetry(err error) bool {
	errs, ok := err.(*Error)
	if ok {
		return errs.IsShouldRetry()
	}
	return false
}
