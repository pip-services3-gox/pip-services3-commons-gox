package test_data

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

func TestAnyValueMapNew(t *testing.T) {
	mp := data.NewEmptyAnyValueMap()
	_, ok := mp.GetAsObject("key1")
	assert.False(t, ok)

	mp = data.NewAnyValueMapFromValue(map[string]interface{}{
		"key1": 1,
		"key2": "A",
	})
	val, ok := mp.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, int64(1), val)

	val, ok = mp.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)

	mp = data.NewAnyValueMapFromMaps(map[string]interface{}{
		"key1": 1,
		"key2": "A",
	})
	val, ok = mp.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = mp.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)

	mp = data.NewAnyValueMapFromTuples(
		"key1", 1,
		"key2", "A",
	)
	val, ok = mp.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	val, ok = mp.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "A", val)
}

func TestAnyValueMapGetAndSet(t *testing.T) {
	mp := data.NewEmptyAnyValueMap()
	_, ok := mp.GetAsObject("key1")
	assert.False(t, ok)

	mp.SetAsObject("key1", 1)
	assert.Equal(t, 1, mp.GetAsInteger("key1"))
	assert.True(t, 1.0-mp.GetAsFloat("key1") < 0.001)
	assert.Equal(t, "1", mp.GetAsString("key1"))

	mp.Put("key2", "1")
	assert.Equal(t, 1, mp.GetAsInteger("key2"))
	assert.True(t, 1.0-mp.GetAsFloat("key2") < 0.001)
	assert.Equal(t, "1", mp.GetAsString("key2"))

	mp.Remove("key2")
	_, ok = mp.GetAsObject("key2")
	assert.False(t, ok)
}
