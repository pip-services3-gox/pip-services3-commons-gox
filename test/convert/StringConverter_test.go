package test_convert

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	str, ok := convert.StringConverter.ToNullableString(nil)
	assert.False(t, ok)
	assert.Equal(t, "", str)

	assert.Equal(t, "xyz", convert.StringConverter.ToString("xyz"))
	assert.Equal(t, "123", convert.StringConverter.ToString(123))
	assert.Equal(t, "true", convert.StringConverter.ToString(true))

	value := struct{ prop string }{"xyz"}
	assert.Equal(t, "{xyz}", convert.StringConverter.ToString(value))

	array1 := []string{"A", "B", "C"}
	assert.Equal(t, "A,B,C", convert.StringConverter.ToString(array1))

	array2 := []int32{1, 2, 3}
	assert.Equal(t, "1,2,3", convert.StringConverter.ToString(array2))

	assert.Equal(t, "xyz", convert.StringConverter.ToStringWithDefault(nil, "xyz"))
}
