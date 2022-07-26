package validator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cymonevo-luna/go-core/pkg/errors"
)

type Validator interface {
	Missing(param string)
	Message(message string)
	Error() error
}

type validatorImpl struct {
	missing []string
	errs    []string
}

// New validator for unified error validation
func New() Validator {
	return &validatorImpl{}
}

// Missing adds if there's missing param violation
// can be added multiple times
func (v *validatorImpl) Missing(param string) {
	v.missing = append(v.missing, param)
}

// Message adds if there's custom violation
// can be added multiple times
func (v *validatorImpl) Message(message string) {
	v.errs = append(v.errs, message)
}

// Error build the error output based on all violations
// return nil if there's no violations
func (v *validatorImpl) Error() error {
	if len(v.missing) == 0 && len(v.errs) == 0 {
		return nil
	}
	var msgs []string
	if len(v.missing) > 0 {
		msgs = append(msgs, fmt.Sprintf("Missing Param(s) [%s]", strings.Join(v.missing, ", ")))
	}
	if len(v.errs) > 0 {
		msgs = append(msgs, v.errs...)
	}
	return errors.New(http.StatusUnprocessableEntity).
		WithMessage(strings.Join(msgs, "; "))
}
