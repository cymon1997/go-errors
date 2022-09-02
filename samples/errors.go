package samples

import (
	"log"
	"net/http"

	"github.com/cymon1997/go-errors/errors"
	"github.com/cymon1997/go-errors/validator"
)

func samples() {
	// Define basic errors
	var (
		ErrInvalidRequest = errors.New(http.StatusUnprocessableEntity).
					WithMessage("Invalid Request")
		ErrInternalServer = errors.New(http.StatusInternalServerError).
					WithMessage("Internal Server")
	)

	// Status code retrieval
	errors.GetStatus(ErrInvalidRequest)
	errors.GetMessage(ErrInvalidRequest)
	// output: status=422 message="Invalid Request"

	// Check if error should be retried
	errors.IsShouldRetry(ErrInvalidRequest)
	// output: false
	errors.IsShouldRetry(ErrInternalServer)
	// output: true

	type Input struct {
		Mandatory interface{} `json:"mandatory"`
		Positive  int         `json:"positive"`
	}

	// Sample validator in function
	fn := func(input Input) error {
		v := validator.New()
		if input.Mandatory == nil {
			v.Missing("mandatory")
		}
		if input.Positive < 0 {
			v.Message("positive should not be negative")
		}
		return v.Error()
	}

	// Sample error message builder
	err := fn(Input{
		Mandatory: nil,
		Positive:  -1,
	})
	if err != nil {
		log.Print(err.Error())
	}
	// output:
}
