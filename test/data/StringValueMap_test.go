package test_data

import (
	"encoding/json"
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

func TestStringValueMapNew(t *testing.T) {
	value := data.NewEmptyStringValueMap()
	obj, ok := value.GetAsObject("key1")
	assert.False(t, ok)
	assert.Equal(t, "", obj)

	value = data.NewStringValueMapFromValue(map[string]string{
		"key1": "1",
		"key2": "A",
	})
	val, ok := value.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "1", val)
	val, ok = value.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)

	value = data.NewStringValueMapFromMaps(map[string]string{
		"key1": "1",
		"key2": "A",
	})
	val, ok = value.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "1", val)
	val, ok = value.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)

	value = data.NewStringValueMapFromTuples(
		"key1", "1",
		"key2", "A",
	)
	val, ok = value.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "1", val)
	val, ok = value.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)
}

func TestStringValueMapGetAndSet(t *testing.T) {
	value := data.NewEmptyStringValueMap()
	obj, ok := value.GetAsObject("key1")
	assert.False(t, ok)
	assert.Equal(t, "", obj)

	value.SetAsObject("key1", 1)
	assert.Equal(t, 1, value.GetAsInteger("key1"))
	assert.True(t, 1.0-value.GetAsFloat("key1") < 0.001)
	assert.Equal(t, "1", value.GetAsString("key1"))

	value.Put("key2", "1")
	assert.Equal(t, 1, value.GetAsInteger("key2"))
	assert.True(t, 1.0-value.GetAsFloat("key2") < 0.001)
	assert.Equal(t, "1", value.GetAsString("key2"))

	value.Remove("key2")
	obj, ok = value.GetAsObject("key2")
	assert.False(t, ok)
	assert.Equal(t, "", obj)
}

func TestStringValueMapJsonSerialization(t *testing.T) {
	json1 := []byte("{\"key1\":\"1\",\"key2\":\"A\"}")

	var value *data.StringValueMap
	err := json.Unmarshal(json1, &value)
	assert.Empty(t, err)
	val, ok := value.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "1", val)
	val, ok = value.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)

	json2, err2 := json.Marshal(value)
	assert.Empty(t, err2)
	assert.Equal(t, json1, json2)
}
