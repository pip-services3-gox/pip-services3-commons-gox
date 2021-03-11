package test_validate

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	"github.com/stretchr/testify/assert"
)

func TestOrRule(t *testing.T) {
	obj := &TestClass{}

	schema := validate.NewSchema().
		WithRule(validate.NewOrRule(
			validate.NewAtLeastOneExistsRule("missingProperty", "stringField1", "nullProperty"),
			validate.NewAtLeastOneExistsRule("stringField1", "nullProperty", "intField"),
		))
	results := schema.Validate(obj)
	assert.Equal(t, 0, len(results))

	schema = validate.NewSchema().
		WithRule(validate.NewOrRule(
			validate.NewAtLeastOneExistsRule("missingProperty", "stringField1", "nullProperty"),
			validate.NewAtLeastOneExistsRule("missingProperty", "nullProperty"),
		))
	results = schema.Validate(obj)
	assert.Equal(t, 0, len(results))
}
