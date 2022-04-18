package test_convert

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/stretchr/testify/assert"
)

func TestJsonToMap(t *testing.T) {
	// Handling simple objects
	v := `{ "value1":123, "value2":234 }`
	m := convert.JsonConverter.ToMap(v)
	assert.NotNil(t, m)
	assert.Len(t, m, 2)
	assert.Equal(t, 123., m["value1"])
	assert.Equal(t, 234., m["value2"])

	// Recursive conversion
	v = `{ "value1":123, "value2": { "value21": 111, "value22": 222} }`
	m = convert.JsonConverter.ToMap(v)
	assert.NotNil(t, m)
	assert.Len(t, m, 2)
	assert.Equal(t, 123., m["value1"])

	m2 := m["value2"].(map[string]interface{})
	assert.Len(t, m2, 2)
	assert.Equal(t, 111., m2["value21"])
	assert.Equal(t, 222., m2["value22"])

	// Handling arrays
	v = `{ "value1":123, "value2": [{ "value21": 111, "value22": 222}] }`
	m = convert.JsonConverter.ToMap(v)
	assert.NotNil(t, m)
	assert.Len(t, m, 2)
	assert.Equal(t, 123., m["value1"])

	a2 := m["value2"].([]interface{})
	assert.NotNil(t, a2)
	assert.Len(t, a2, 1)

	m2 = a2[0].(map[string]interface{})
	assert.Len(t, m2, 2)
	assert.Equal(t, 111., m2["value21"])
	assert.Equal(t, 222., m2["value22"])
}
