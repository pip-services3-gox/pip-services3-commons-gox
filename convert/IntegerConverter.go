package convert

// IntegerConverter converts arbitrary values into integer using extended conversion rules:
// - Strings are converted to integer values
// - DateTime: total number of milliseconds since unix epoсh
// - Boolean: 1 for true and 0 for false
//
// Example:
//
//  value1, ok1 := convert.IntegerConverter.ToNullableInteger("ABC")
//  value2, ok2 := convert.IntegerConverter.ToNullableInteger("123.456")
//  value3, ok3 := convert.IntegerConverter.ToNullableInteger(true)
//  value4, ok4 := convert.IntegerConverter.ToNullableInteger(time.Now())
//  fmt.Println(value1, ok1) // 0, false
//  fmt.Println(value2, ok2) // 123, true
//  fmt.Println(value3, ok3) // 1, true
//  fmt.Println(value4, ok4) // current milliseconds (e.g. 1566333428), true
var IntegerConverter = &_TIntegerConverter{}

type _TIntegerConverter struct{}

// ToNullableInteger converts value into integer or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value and true or 0 and false when conversion is not supported.
func (c *_TIntegerConverter) ToNullableInteger(value any) (int, bool) {
	return toNullableInteger(value)
}

// ToInteger converts value into integer or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or 0 when conversion is not supported.
func (c *_TIntegerConverter) ToInteger(value any) int {
	return toInteger(value)
}

// ToIntegerWithDefault converts value into integer or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: integer value or default when conversion is not supported.
func (c *_TIntegerConverter) ToIntegerWithDefault(value any, defaultValue int) int {
	return toIntegerWithDefault(value, defaultValue)
}

// ToNullableInteger converts value into integer or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or null when conversion is not supported.
func toNullableInteger(value any) (int, bool) {
	if v, ok := LongConverter.ToNullableLong(value); ok {
		return int(v), ok
	}
	return 0, false
}

// ToInteger converts value into integer or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or 0 when conversion is not supported.
func toInteger(value any) int {
	return toIntegerWithDefault(value, 0)
}

// ToIntegerWithDefault converts value into integer or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: integer value or default when conversion is not supported.
func toIntegerWithDefault(value any, defaultValue int) int {
	if r, ok := toNullableInteger(value); ok {
		return r
	}
	return defaultValue
}
