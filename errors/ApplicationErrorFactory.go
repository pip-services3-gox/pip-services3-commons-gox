package errors

/*
Factory to recreate exceptions from ErrorDescription values passed through the wire.
see
ErrorDescription
see
ApplicationError
*/
type TApplicationErrorFactory struct{}

var ApplicationErrorFactory *TApplicationErrorFactory = &TApplicationErrorFactory{}

// Recreates ApplicationError object from serialized ErrorDescription.
// It tries to restore original exception type using type or error category fields.
// Parameters:
//  - description: ErrorDescription
//  a serialized error description received as a result of remote call

// Returns *ApplicationError
func (c *TApplicationErrorFactory) Create(description *ErrorDescription) *ApplicationError {
	return NewErrorFromDescription(description)
}

// Recreates ApplicationError object from description.
// It tries to restore original exception type using type or error category fields.
// Parameters:
//  - description: ErrorDescription
//  a serialized error description received as a result of remote call

// Returns *ApplicationError
func NewErrorFromDescription(description *ErrorDescription) *ApplicationError {
	if description == nil {
		return nil
	}

	var err *ApplicationError = nil
	category := description.Category
	code := description.Code
	message := description.Message
	correlationId := description.CorrelationId

	// Create well-known exception type based on error category
	if Unknown == category {
		err = NewUnknownError(correlationId, code, message)
	} else if Internal == category {
		err = NewInternalError(correlationId, code, message)
	} else if Misconfiguration == category {
		err = NewConfigError(correlationId, code, message)
	} else if NoResponse == category {
		err = NewConnectionError(correlationId, code, message)
	} else if FailedInvocation == category {
		err = NewInvocationError(correlationId, code, message)
	} else if FileError == category {
		err = NewFileError(correlationId, code, message)
	} else if BadRequest == category {
		err = NewBadRequestError(correlationId, code, message)
	} else if Unauthorized == category {
		err = NewUnauthorizedError(correlationId, code, message)
	} else if Conflict == category {
		err = NewConflictError(correlationId, code, message)
	} else if NotFound == category {
		err = NewNotFoundError(correlationId, code, message)
	} else if InvalidState == category {
		err = NewInvalidStateError(correlationId, code, message)
	} else if Unsupported == category {
		err = NewUnsupportedError(correlationId, code, message)
	} else {
		err = NewUnknownError(correlationId, code, message)
		err.Category = category
		err.Status = description.Status
	}

	// Fill error with details
	err.Details = description.Details
	err.Cause = description.Cause
	err.StackTrace = description.StackTrace

	return err
}
