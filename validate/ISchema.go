package validate

import "github.com/pip-services3-gox/pip-services3-commons-gox/errors"

// ISchema validation schema interface
type ISchema interface {
	Validate(value any) []*ValidationResult
	ValidateAndReturnError(correlationId string, value any, strict bool) *errors.ApplicationError
	ValidateAndThrowError(correlationId string, value any, strict bool)
}
