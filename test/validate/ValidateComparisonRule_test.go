package test_validate

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	"github.com/stretchr/testify/assert"
)

func TestValidateComparisonRule(t *testing.T) {
	schema := validate.NewSchema().
		WithRule(validate.NewValueComparisonRule("EQ", 123))
	results := schema.Validate(123)
	assert.Equal(t, 0, len(results))

	results = schema.Validate(423)
	assert.Equal(t, 1, len(results))
}
