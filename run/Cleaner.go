package run

// Cleaner helper class that cleans stored object state.
var Cleaner = &_TCleaner{}

type _TCleaner struct{}

// ClearOne clears state of specific component.
// To be cleaned state components must implement ICleanable interface. If they don't the call to this method has no effect.
//	Parameters:
//		- correlationId: string transaction id to trace execution through call chain.
//		- component any the component that is to be cleaned.
//	Returns: error
func (c *_TCleaner) ClearOne(correlationId string, component any) error {
	if v, ok := component.(ICleanable); ok {
		return v.Clear(correlationId)
	}
	return nil
}

// Clear clears state of multiple components.
// To be cleaned state components must implement ICleanable interface.
// If they don't the call to this method has no effect.
//	Parameters:
//		- correlationId string transaction id to trace execution through call chain.
//		- components []any the list of components that are to be cleaned.
//	Returns: error
func (c *_TCleaner) Clear(correlationId string, components []any) error {
	for _, component := range components {
		err := c.ClearOne(correlationId, component)
		if err != nil {
			return err
		}
	}
	return nil
}
