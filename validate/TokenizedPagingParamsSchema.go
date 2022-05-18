package validate

//Schema to validate TokenizedPagingParams.

import "github.com/pip-services3-gox/pip-services3-commons-gox/convert"

// NewTokenizedPagingParamsSchema creates a new instance of validation schema.
//	Returns: *TokenizedPagingParamsSchema
func NewTokenizedPagingParamsSchema() *ObjectSchema {
	return NewObjectSchema().
		WithOptionalProperty("token", convert.String).
		WithOptionalProperty("take", convert.Long).
		WithOptionalProperty("total", convert.Boolean)
}
