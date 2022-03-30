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
//  value1 := convert.LongConverter.ToNullableLong("ABC")
//  value2 := convert.LongConverter.ToNullableLong("123.456")
//  value3 := convert.LongConverter.ToNullableLong(true)
//  value4 := convert.LongConverter.ToNullableLong(time.Now())
//  fmt.Println(value1)  // <nil>
//  fmt.Println(*value2) // 123
//  fmt.Println(*value3) // 1
//  fmt.Println(*value4) // current milliseconds (e.g. 1566333527)
var LongConverter = &_TLongConverter{}

type _TLongConverter struct{}

// ToNullableLong converts value into long or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or null when conversion is not supported.
func (c *_TLongConverter) ToNullableLong(value any) *int64 {
	return ToNullableLong(value)
}

// ToLong converts value into long or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or 0 when conversion is not supported.
func (c *_TLongConverter) ToLong(value any) int64 {
	return ToLong(value)
}

// ToLongWithDefault converts value into long or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value..
// Returns: long value or default when conversion is not supported.
func (c *_TLongConverter) ToLongWithDefault(value any, defaultValue int64) int64 {
	return ToLongWithDefault(value, defaultValue)
}

// ToNullableLong converts value into long or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or null when conversion is not supported.
func ToNullableLong(value any) *int64 {
	if value == nil {
		return nil
	}

	var r int64 = 0

	switch value.(type) {
	case int8:
		r = (int64)(value.(int8))
	case uint8:
		r = (int64)(value.(uint8))
	case int:
		r = (int64)(value.(int))
	case int16:
		r = (int64)(value.(int16))
	case uint16:
		r = (int64)(value.(uint16))
	case int32:
		r = (int64)(value.(int32))
	case uint32:
		r = (int64)(value.(uint32))
	case int64:
		r = (int64)(value.(int64))
	case uint64:
		r = (int64)(value.(uint64))
	case float32:
		r = (int64)(value.(float32))
	case float64:
		r = (int64)(value.(float64))

	case bool:
		v := value.(bool)
		if v == true {
			r = 1
		}

	case time.Time:
		r = value.(time.Time).Unix()

	case time.Duration:
		r = value.(time.Duration).Nanoseconds() / 1000000

	case string:
		v, ok := strconv.ParseFloat(value.(string), 0)
		if ok != nil {
			return nil
		}
		r = int64(v)

	default:
		return nil
	}

	return &r
}

// ToLong converts value into long or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: long value or 0 when conversion is not supported.
func ToLong(value any) int64 {
	return ToLongWithDefault(value, 0)
}

// ToLongWithDefault converts value into long or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value..
// Returns: long value or default when conversion is not supported.
func ToLongWithDefault(value any, defaultValue int64) int64 {
	r := ToNullableLong(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
