package validator

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/cymon1997/go-errors/errors"
)

type Validator interface {
	Required(val interface{}, name string)
	Positive(val int, name string)
	Negative(val int, name string)
	Error() error
}

type validatorImpl struct {
	required  []string
	positives []string
	negatives []string
	errs      []string
}

// New validator for unified error validation
func New() Validator {
	return &validatorImpl{}
}

// Required an input param val as required
func (v *validatorImpl) Required(val interface{}, name string) {
	temp := reflect.ValueOf(val)
	if val == nil || temp.IsZero() || temp.Len() == 0 || temp.Type().Len() == 0 {
		v.required = append(v.required, actualMessage(val, name))
		return
	}
}

// Positive force an input param val as positive number (zero inclusive)
func (v *validatorImpl) Positive(val int, name string) {
	if val < 0 {
		v.positives = append(v.positives, actualMessage(val, name))
	}
}

// Negative force an input param val as negative number (zero exclusive)
func (v *validatorImpl) Negative(val int, name string) {
	if val >= 0 {
		v.negatives = append(v.negatives, actualMessage(val, name))
	}
}

// Custom adds if there's custom violation
func (v *validatorImpl) Custom(val interface{}, msg string) {
	v.errs = append(v.errs, actualMessage(val, msg))
}

// Error build the error output based on all violations
// return nil if there's no violations
func (v *validatorImpl) Error() error {
	if len(v.required) == 0 && len(v.positives) == 0 && len(v.negatives) == 0 && len(v.errs) == 0 {
		return nil
	}
	var msgs []string
	if len(v.required) > 0 {
		msgs = append(msgs, fmt.Sprintf("Required Param(s) [%s]", strings.Join(v.required, ", ")))
	}
	if len(v.positives) > 0 {
		msgs = append(msgs, fmt.Sprintf("Positive Param(s) [%s]", strings.Join(v.positives, ", ")))
	}
	if len(v.negatives) > 0 {
		msgs = append(msgs, fmt.Sprintf("Negative Param(s) [%s]", strings.Join(v.negatives, ", ")))
	}
	if len(v.errs) > 0 {
		msgs = append(msgs, v.errs...)
	}
	return errors.New(http.StatusUnprocessableEntity).
		WithMessage(strings.Join(msgs, "; "))
}
