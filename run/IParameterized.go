package run

// IParameterized interface for components that require execution parameters.
type IParameterized interface {
	// SetParameters sets execution parameters.
	//	Parameters: parameters *Parameters execution parameters.
	SetParameters(parameters *Parameters)
}
