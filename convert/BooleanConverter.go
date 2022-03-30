package convert

import (
	"fmt"
	"strings"
	"time"
)

// BooleanConverter converts arbitrary values to boolean values using extended conversion rules:
// 	- Numbers: above 0, less much 0 are true; equal to 0 are false
// 	- Strings: "true", "yes", "T", "Y", "1" are true, "false", "no", "F", "N" are false
// 	- DateTime: above 0, less much 0 total milliseconds are true, equal to 0 are false
//
// Example:
//
//  value1 := convert.BooleanConverter.ToNullableBoolean(true)
//  value2 := convert.BooleanConverter.ToNullableBoolean("yes")
//  value3 := convert.BooleanConverter.ToNullableBoolean(1)
//  value4 := convert.BooleanConverter.ToNullableBoolean(struct{}{})
//  fmt.Println(*value1) // true
//  fmt.Println(*value2) // true
//  fmt.Println(*value3) // true
//  fmt.Println(value4)  // <nil>
var BooleanConverter = &_TBooleanConverter{}

type _TBooleanConverter struct{}

// ToNullableBoolean converts value into boolean or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: boolean value or null when conversion is not supported.
func (c *_TBooleanConverter) ToNullableBoolean(value any) *bool {
	return ToNullableBoolean(value)
}

// ToBoolean converts value into boolean or returns false when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: boolean value or false when conversion is not supported.
func (c *_TBooleanConverter) ToBoolean(value any) bool {
	return ToBoolean(value)
}

// ToBooleanWithDefault converts value into boolean or returns default value when conversion is not possible
// Parameters:"value" - the value to convert.
//  "defaultValue" - the default value
// Returns: boolean value or default when conversion is not supported.
func (c *_TBooleanConverter) ToBooleanWithDefault(value any, defaultValue bool) bool {
	return ToBooleanWithDefault(value, defaultValue)
}

// ToNullableBoolean converts value into boolean or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: boolean value or null when conversion is not supported.
func ToNullableBoolean(value any) *bool {
	if value == nil {
		return nil
	}

	var v string

	switch value.(type) {
	case bool:
		r := value.(bool)
		return &r

	case string:
		v = strings.ToLower(value.(string))

	case time.Duration:
		d := value.(time.Duration)
		r := d.Nanoseconds() > 0
		return &r

	default:
		v = strings.ToLower(fmt.Sprint(value))
	}

	if v == "1" || v == "true" || v == "t" || v == "yes" || v == "y" {
		r := true
		return &r
	}

	if v == "0" || v == "false" || v == "f" || v == "no" || v == "n" {
		r := false
		return &r
	}

	return nil
}

// ToBoolean converts value into boolean or returns false when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: boolean value or false when conversion is not supported.
func ToBoolean(value any) bool {
	return ToBooleanWithDefault(value, false)
}

// ToBooleanWithDefault converts value into boolean or returns default value when conversion is not possible
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: boolean value or default when conversion is not supported.
func ToBooleanWithDefault(value any, defaultValue bool) bool {
	r := ToNullableBoolean(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
