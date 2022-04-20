package test_validate

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	"github.com/stretchr/testify/assert"
)

func TestArraySchema(t *testing.T) {
	schema := validate.NewObjectSchema().
		WithRequiredProperty("intField", convert.Integer).
		WithRequiredProperty("stringField1", convert.String).
		WithRequiredProperty("stringField2", convert.String).
		WithRequiredProperty("intArrayField", validate.NewArraySchema(convert.Integer)).
		WithRequiredProperty("stringArrayField", validate.NewArraySchema(convert.String)).
		WithRequiredProperty("mapField", convert.Map).
		WithRequiredProperty("subObjectField", convert.Object).
		WithRequiredProperty("subArrayField", validate.NewArraySchema(convert.Object))

	obj := &TestClass{
		IntArrayField:    []int{1, 2, 3},
		StringArrayField: []string{"A", "B", "C"},
		MapField:         map[string]any{},
		SubObjectField:   &SubTestClass{},
		SubArrayField:    []*SubTestClass{},
	}
	results := schema.Validate(obj)
	assert.Equal(t, 0, len(results))
}
