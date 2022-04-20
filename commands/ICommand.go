package commands

import (
	"github.com/pip-services3-gox/pip-services3-commons-gox/run"
	"github.com/pip-services3-gox/pip-services3-commons-gox/validate"
)

// ICommand An interface for Commands, which are part of the Command design pattern.
// Each command wraps a method or function and allows to call them in uniform and safe manner.
type ICommand interface {
	run.IExecutable
	// Name gets the command name.
	//	Returns: string the command name.
	Name() string
	// Validate validates command arguments before execution using defined schema.
	//	see Parameters
	//	see ValidationResult
	//	Parameters: args: Parameters the parameters (arguments) to validate.
	//	Returns: ValidationResult[] an array of ValidationResults.
	Validate(args *run.Parameters) []*validate.ValidationResult
}
