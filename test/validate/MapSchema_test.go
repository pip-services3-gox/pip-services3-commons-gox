package test_validate

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/validate"
	"github.com/stretchr/testify/assert"
)

func TestMapSchema(t *testing.T) {
	schema := validate.NewObjectSchema().
		WithRequiredProperty("intField", convert.Integer).
		WithRequiredProperty("stringField1", convert.String).
		WithRequiredProperty("stringField2", convert.String).
		WithRequiredProperty("intArrayField", convert.Array).
		WithRequiredProperty("stringArrayField", convert.Array).
		WithRequiredProperty("mapField", validate.NewMapSchema(convert.String, convert.Object)).
		WithRequiredProperty("subObjectField", convert.Object).
		WithRequiredProperty("subArrayField", convert.Array)

	obj := &TestClass{
		IntArrayField:    []int{1, 2, 3},
		StringArrayField: []string{"A", "B", "C"},
		MapField:         map[string]interface{}{},
		SubObjectField:   &SubTestClass{},
		SubArrayField:    []*SubTestClass{},
	}
	results := schema.Validate(obj)
	assert.Equal(t, 0, len(results))
}
