package convert

import (
	"strconv"
	"time"
)

// LongConverter Converts arbitrary values into long using extended conversion rules:
// - Strings are converted to long values
// - DateTime: total number of milliseconds since unix epoсh
// - Boolean: 1 for true and 0 for false
//
// Example:
//
//  value1, ok1 := convert.LongConverter.ToNullableLong("ABC")
//  value2, ok2 := convert.LongConverter.ToNullableLong("123.456")
//  value3, ok3 := convert.LongConverter.ToNullableLong(true)
//  value4, ok4 := convert.LongConverter.ToNullableLong(time.Now())
//  fmt.Println(value1, ok1) // 0, false
//  fmt.Println(value2, ok2) // 123, false
//  fmt.Println(value3, ok3) // 1, false
//  fmt.Println(value4, ok4) // current milliseconds (e.g. 1566333527), false
var LongConverter = &_TLongConverter{}

type _TLongConverter struct{}

// ToNullableLong converts value into long or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or null when conversion is not supported.
func (c *_TLongConverter) ToNullableLong(value any) (int64, bool) {
	return toNullableLong(value)
}

// ToLong converts value into long or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or 0 when conversion is not supported.
func (c *_TLongConverter) ToLong(value any) int64 {
	return toLong(value)
}

// ToLongWithDefault converts value into long or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value..
// Returns: long value or default when conversion is not supported.
func (c *_TLongConverter) ToLongWithDefault(value any, defaultValue int64) int64 {
	return toLongWithDefault(value, defaultValue)
}

// ToNullableLong converts value into long or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value and true or 0 and false when conversion is not supported.
func toNullableLong(value any) (int64, bool) {
	if value == nil {
		return 0, false
	}

	switch value.(type) {
	case int8:
		r, ok := value.(int8)
		return int64(r), ok
	case uint8:
		r, ok := value.(uint8)
		return int64(r), ok
	case int:
		r, ok := value.(int)
		return int64(r), ok
	case int16:
		r, ok := value.(int16)
		return int64(r), ok
	case uint16:
		r, ok := value.(uint16)
		return int64(r), ok
	case int32:
		r, ok := value.(int32)
		return int64(r), ok
	case uint32:
		r, ok := value.(uint32)
		return int64(r), ok
	case int64:
		r, ok := value.(int64)
		return r, ok
	case uint64:
		r, ok := value.(uint64)
		return int64(r), ok
	case float32:
		r, ok := value.(float32)
		return int64(r), ok
	case float64:
		r, ok := value.(float64)
		return int64(r), ok

	case bool:
		if r, ok := value.(bool); ok {
			if r {
				return 1, true
			} else {
				return 0, true
			}
		}
		break

	case time.Time:
		if r, ok := value.(time.Time); ok {
			return r.Unix(), true
		}
		break

	case time.Duration:
		if r, ok := value.(time.Duration); ok {
			return r.Nanoseconds() / 1000000, true
		}
		break

	case string:
		if r, ok := value.(string); ok {
			if v, err := strconv.ParseFloat(r, 0); err == nil {
				return int64(v), true
			}
		}
		break
	}

	return 0, false
}

// ToLong converts value into long or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or 0 when conversion is not supported.
func toLong(value any) int64 {
	return toLongWithDefault(value, 0)
}

// ToLongWithDefault converts value into long or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value..
// Returns: long value or default when conversion is not supported.
func toLongWithDefault(value any, defaultValue int64) int64 {
	if r, ok := toNullableLong(value); ok {
		return r
	}
	return defaultValue
}
