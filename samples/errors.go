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
	errors.IsRetry(ErrInvalidRequest)
	// output: false
	errors.IsRetry(ErrInternalServer)
	// output: true

	type Input struct {
		Mandatory interface{} `json:"a"`
		Positive  int         `json:"b"`
	}

	// Sample validator in function
	fn := func(input Input) error {
		v := validator.New()
		v.Required(input.Mandatory, "a")
		v.Positive(input.Positive, "b")
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
