package convert

// FloatConverter Converts arbitrary values into float using extended conversion rules:
// - Strings are converted to float values
// - DateTime: total number of milliseconds since unix epoсh
// - Boolean: 1 for true and 0 for false
//
// Example:
//
//  value1 := convert.FloatConverter.ToNullableFloat("ABC")
//  value2 := convert.FloatConverter.ToNullableFloat("123.456")
//  value3 := convert.FloatConverter.ToNullableFloat(true)
//  value4 := convert.FloatConverter.ToNullableFloat(time.Now())
//  fmt.Println(value1)  // <nil>
//  fmt.Println(*value2) // 123.456
//  fmt.Println(*value3) // 1
//  fmt.Println(*value4) // current milliseconds (e.g. 1.566333114e+09)
var FloatConverter = &_TFloatConverter{}

type _TFloatConverter struct{}

// ToNullableFloat converts value into float or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: float value or null when conversion is not supported.
func (c *_TFloatConverter) ToNullableFloat(value any) *float32 {
	return ToNullableFloat(value)
}

// ToFloat converts value into float or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: float value or 0 when conversion is not supported.
func (c *_TFloatConverter) ToFloat(value any) float32 {
	return ToFloat(value)
}

// ToFloatWithDefault converts value into float or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: float value or default when conversion is not supported.
func (c *_TFloatConverter) ToFloatWithDefault(value any, defaultValue float32) float32 {
	return ToFloatWithDefault(value, defaultValue)
}

// ToNullableFloat converts value into float or returns null when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: float value or null when conversion is not supported.
func ToNullableFloat(value any) *float32 {
	v := ToNullableDouble(value)
	if v == nil {
		return nil
	}
	r := float32(*v)
	return &r
}

// ToFloat converts value into float or returns 0 when conversion is not possible.
// Parameters: "value" - the value to convert.
// Returns: float value or 0 when conversion is not supported.
func ToFloat(value any) float32 {
	return ToFloatWithDefault(value, 0)
}

// ToFloatWithDefault converts value into float or returns default when conversion is not possible.
// Parameters:
//  "value" - the value to convert.
//  "defaultValue" - the default value.
// Returns: float value or default when conversion is not supported.
func ToFloatWithDefault(value any, defaultValue float32) float32 {
	r := ToNullableFloat(value)
	if r == nil {
		return defaultValue
	}
	return *r
}
