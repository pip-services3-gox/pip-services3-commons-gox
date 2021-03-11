package test_validate

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	"github.com/stretchr/testify/assert"
)

func TestNotRule(t *testing.T) {
	obj := &TestClass{}

	schema := validate.NewSchema().
		WithRule(validate.NewNotRule(
			validate.NewAtLeastOneExistsRule("stringField1", "nullProperty", "intField"),
		))
	results := schema.Validate(obj)
	assert.Equal(t, 1, len(results))

	schema = validate.NewSchema().
		WithRule(validate.NewNotRule(
			validate.NewAtLeastOneExistsRule("missingProperty", "nullProperty"),
		))
	results = schema.Validate(obj)
	assert.Equal(t, 0, len(results))
}
