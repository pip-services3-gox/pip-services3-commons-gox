package validate

/*
Schema to validate FilterParams.
*/
import "github.com/pip-services3-gox/pip-services3-commons-gox/convert"

// Creates a new instance of validation schema.
// Returns *MapSchema
func NewFilterParamsSchema() *MapSchema {
	return NewMapSchema(convert.String, nil)
}
