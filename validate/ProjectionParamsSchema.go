package validate

// Schema to validate ProjectionParams

import "github.com/pip-services3-gox/pip-services3-commons-gox/convert"

// NewProjectionParamsSchema creates a new instance of validation schema.
//	Returns: *ArraySchema
func NewProjectionParamsSchema() *ArraySchema {
	return NewArraySchema(convert.String)
}
