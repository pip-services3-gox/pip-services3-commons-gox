package convert

// IntegerConverter converts arbitrary values into integer using extended conversion rules:
// - Strings are converted to integer values
// - DateTime: total number of milliseconds since unix epoсh
// - Boolean: 1 for true and 0 for false
//
// Example:
//
//  value1 := convert.IntegerConverter.ToNullableInteger("ABC")
//  value2 := convert.IntegerConverter.ToNullableInteger("123.456")
//  value3 := convert.IntegerConverter.ToNullableInteger(true)
//  value4 := convert.IntegerConverter.ToNullableInteger(time.Now())
//  fmt.Println(value1)  // <nil>
//  fmt.Println(*value2) // 123
//  fmt.Println(*value3) // 1
//  fmt.Println(*value4) // current milliseconds (e.g. 1566333428)
var IntegerConverter = &_TIntegerConverter{}

type _TIntegerConverter struct{}

// ToNullableInteger converts value into integer or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or null when conversion is not supported.
func (c *_TIntegerConverter) ToNullableInteger(value any) *int {
	return ToNullableInteger(value)
}

// ToInteger converts value into integer or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or 0 when conversion is not supported.
func (c *_TIntegerConverter) ToInteger(value any) int {
	return ToInteger(value)
}

// ToIntegerWithDefault converts value into integer or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: integer value or default when conversion is not supported.
func (c *_TIntegerConverter) ToIntegerWithDefault(value any, defaultValue int) int {
	return ToIntegerWithDefault(value, defaultValue)
}

// ToNullableInteger converts value into integer or returns null when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or null when conversion is not supported.
func ToNullableInteger(value any) *int {
	v := ToNullableLong(value)
	if v == nil {
		return nil
	}
	r := int(*v)
	return &r
}

// ToInteger converts value into integer or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert
// Returns: integer value or 0 when conversion is not supported.
func ToInteger(value any) int {
	return ToIntegerWithDefault(value, 0)
}

// ToIntegerWithDefault converts value into integer or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: integer value or default when conversion is not supported.
func ToIntegerWithDefault(value any, defaultValue int) int {
	r := ToNullableInteger(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
