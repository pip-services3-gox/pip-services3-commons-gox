package run

import "context"

// ContextFeedbackChan a channel to send
// default context feedback
type ContextFeedbackChan chan int8

// ContextFeedbackChanWithError a channel to send
// context feedback with error
type ContextFeedbackChanWithError chan error

// ContextFeedbackChanWithCustomData a channel to
// send context feedback with specific data.
type ContextFeedbackChanWithCustomData[T any] chan T

// ContextValueType an enum to describe specific
// context feedback channel
//	Possible values:
//		- ContextFeedbackChanType
//		- ContextFeedbackChanWithErrorType
//		- ContextFeedbackChanWithCustomDataType
type ContextValueType string

const (
	ContextFeedbackChanType               ContextValueType = "pip.ContextFeedbackChan"
	ContextFeedbackChanWithErrorType      ContextValueType = "pip.ContextFeedbackChanWithError"
	ContextFeedbackChanWithCustomDataType ContextValueType = "pip.ContextFeedbackChanWithCustomData"
)

const DefaultCancellationSignal int8 = 1

// NewCancelContext wrap context with ContextFeedbackChan
//	see context.WithValue
//	Parameters:
//		- context.Context parent context
//		- ContextFeedbackChan channel to put into context
//	Returns:
//		- context.Context is a context with value
//		- bool true if channel is not nil or false
func NewCancelContext(ctx context.Context, contextFeedbackChan ContextFeedbackChan) (context.Context, bool) {
	if contextFeedbackChan == nil {
		return ctx, false
	}
	return context.WithValue(ctx, ContextFeedbackChanType, contextFeedbackChan), true
}

// NewCancelContextWithError wrap context with ContextFeedbackChanWithError
//	see context.WithValue
//	Parameters:
//		- context.Context - parent context
//		- ContextFeedbackChanWithError - channel to put into context
//	Returns:
//		- context.Context is a context with value
//		- bool true if channel is not nil or false
func NewCancelContextWithError(ctx context.Context, contextFeedbackChan ContextFeedbackChanWithError) (context.Context, bool) {
	if contextFeedbackChan == nil {
		return ctx, false
	}
	return context.WithValue(ctx, ContextFeedbackChanWithErrorType, contextFeedbackChan), true
}

// NewCancelContextWithCustomDataChannel wrap context with ContextFeedbackChanWithCustomData
//	T is a custom data type
//	see context.WithValue
//	Parameters:
//		- context.Context - parent context
//		- ContextFeedbackChanWithCustomData - channel to put into context
//	Returns:
//		- context.Context is a context with value
//		- bool true if channel is not nil or false
func NewCancelContextWithCustomDataChannel[T any](ctx context.Context, contextFeedbackChan ContextFeedbackChanWithCustomData[T]) (context.Context, bool) {
	if contextFeedbackChan == nil {
		return ctx, false
	}
	return context.WithValue(ctx, ContextFeedbackChanWithCustomDataType, contextFeedbackChan), true
}

// CancelContextFeedback sends interrupt signal up to the context owner
//	Parameters: context.Context is a current context
//	Returns: bool true if signal sends successful or false
func CancelContextFeedback(ctx context.Context) bool {
	if val := ctx.Value(ContextFeedbackChanType); val != nil {
		if _chan, ok := val.(ContextFeedbackChan); ok {
			select {
			case _chan <- DefaultCancellationSignal:
				return true
			default:
				return false
			}
		}
	}
	return false
}

// CancelContextFeedbackWithError sends error and interrupt signal up to the context owner
//	Parameters:
//		- context.Context is a current context
//		- error
//	Returns: bool true if signal sends successful or false
func CancelContextFeedbackWithError(ctx context.Context, err error) bool {
	if val := ctx.Value(ContextFeedbackChanWithErrorType); val != nil {
		if _chan, ok := val.(ContextFeedbackChanWithError); ok {
			select {
			case _chan <- err:
				return true
			default:
				return false
			}
		}
	}
	return false
}

// CancelContextFeedbackWithData sends custom data and interrupt signal up to the context owner
//	Parameters:
//		- context.Context is a current context
//		- T custom data
//	Returns: bool true if signal sends successful or false
func CancelContextFeedbackWithData[T any](ctx context.Context, data T) bool {
	if val := ctx.Value(ContextFeedbackChanWithCustomDataType); val != nil {
		if _chan, ok := val.(ContextFeedbackChanWithCustomData[T]); ok {
			select {
			case _chan <- data:
				return true
			default:
				return false
			}
		}
	}
	return false
}
