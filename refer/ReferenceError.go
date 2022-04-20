package refer

// Error when required component dependency cannot be found.

import (
	"fmt"

	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

// NewReferenceError Creates an error instance and assigns its values.
//	Parameters:
//		- correlationId string
//		- locator any the locator to find reference to dependent component.
//	Returns *errors.ApplicationError
func NewReferenceError(correlationId string, locator any) *errors.ApplicationError {
	message := fmt.Sprintf("Failed to obtain reference to %v", locator)
	e := errors.NewInternalError(correlationId, "REF_ERROR", message)
	e.WithDetails("locator", locator)
	return e
}
