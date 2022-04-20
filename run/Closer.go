package run

// Closer helper class that closes previously opened components.
var Closer = &_TCloser{}

type _TCloser struct{}

// Closes specific component.

// CloseOne to be closed components must implement ICloseable interface.
// If they don't the call to this method has no effect.
//	Parameters:
//		- correlationId string transaction id to trace execution through call chain.
//		- component any the component that is to be closed.
//	Returns: error
func (c *_TCloser) CloseOne(correlationId string, component any) error {
	if v, ok := component.(IClosable); ok {
		return v.Close(correlationId)
	}
	return nil
}

// Close closes multiple components.
// To be closed components must implement ICloseable interface.
// If they don't the call to this method has no effect.
//	Parameters:
//		- correlationId string transaction id to trace execution through call chain.
//		- components []any the list of components that are to be closed.
//	Returns: error
func (c *_TCloser) Close(correlationId string, components []any) error {
	for _, component := range components {
		if err := c.CloseOne(correlationId, component); err != nil {
			return err
		}
	}
	return nil
}
