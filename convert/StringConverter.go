package convert

import (
	"fmt"
	refl "reflect"
	"strconv"
	"time"
)

// StringConverter converts arbitrary values into strings using extended conversion rules:
// - Numbers: are converted with '.' as decimal point
// - DateTime: using ISO format
// - Boolean: "true" for true and "false" for false
// - Arrays: as comma-separated list
// - Other objects: using toString() method
//
// Example:
//
//  var value1 = convert.StringConverter.ToString(123.456)
//  var value2 = convert.StringConverter.ToString(true)
//  var value3 = convert.StringConverter.ToString(time.Now())
//  var value4 = convert.StringConverter.ToString([...]int{1, 2, 3})
//  fmt.Println(value1) // 123.456
//  fmt.Println(value2) // true
//  fmt.Println(value3) // 2019-08-20T23:54:47+03:00
//  fmt.Println(value4) // 1,2,3
var StringConverter = &_TStringConverter{}

type _TStringConverter struct{}

// ToNullableString converts value into string or returns null when value is null.
// Parameters: "value" - the value to convert
// Returns: string value or null when value is null.
func (c *_TStringConverter) ToNullableString(value any) *string {
	return ToNullableString(value)
}

// ToString converts value into string or returns "" when value is null.
// Parameters: "value" - the value to convert
// Returns: string value or "" when value is null.
func (c *_TStringConverter) ToString(value any) string {
	return ToString(value)
}

// ToStringWithDefault converts value into string or returns default when value is null.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: string value or default when value is null.
func (c *_TStringConverter) ToStringWithDefault(value any, defaultValue string) string {
	return ToStringWithDefault(value, defaultValue)
}

// ToNullableString converts value into string or returns null when value is null.
// Parameters: "value" - the value to convert
// Returns: string value or null when value is null.
func ToNullableString(value any) *string {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case string:
		r := value.(string)
		return &r

	case byte, uint, uint32, uint64, int, int32, int64:
		r := strconv.FormatInt(ToLong(value), 10)
		return &r

	case float32, float64:
		r := strconv.FormatFloat(ToDouble(value), 'f', -1, 64)
		return &r

	case bool:
		r := "false"
		if value.(bool) {
			r = "true"
		}
		return &r

	case time.Time:
		r := value.(time.Time).Format(time.RFC3339)
		return &r

	case time.Duration:
		r := strconv.FormatInt(value.(time.Duration).Nanoseconds()/1000000, 10)
		return &r

	default:
		val := refl.ValueOf(value)
		if val.Kind() == refl.Slice || val.Kind() == refl.Array {
			r := ""
			for index := 0; index < val.Len(); index++ {
				if len(r) > 0 {
					r += ","
				}
				r += fmt.Sprint(val.Index(index).Interface())
			}
			return &r
		}

		r := fmt.Sprint(value)
		return &r
	}
}

// ToString converts value into string or returns "" when value is null.
// Parameters: "value" - the value to convert
// Returns: string value or "" when value is null.
func ToString(value any) string {
	return ToStringWithDefault(value, "")
}

// ToStringWithDefault converts value into string or returns default when value is null.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: string value or default when value is null.
func ToStringWithDefault(value any, defaultValue string) string {
	r := ToNullableString(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
