package test_convert

import (
	"testing"

	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/stretchr/testify/assert"
)

func TestToLong(t *testing.T) {
	val, ok := convert.LongConverter.ToNullableLong(nil)
	assert.False(t, ok)
	assert.Equal(t, int64(0), val)

	assert.Equal(t, int64(123), convert.LongConverter.ToLong(123))
	assert.Equal(t, int64(123), convert.LongConverter.ToLong(123.456))
	assert.Equal(t, int64(123), convert.LongConverter.ToLong("123"))
	assert.Equal(t, int64(123), convert.LongConverter.ToLong("123.456"))

	assert.Equal(t, int64(123), convert.LongConverter.ToLongWithDefault(nil, 123))
	assert.Equal(t, int64(0), convert.LongConverter.ToLongWithDefault(false, 123))
	assert.Equal(t, int64(123), convert.LongConverter.ToLongWithDefault("ABC", 123))
}
