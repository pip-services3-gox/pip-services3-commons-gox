package convert

import (
	"encoding/json"
)

// JsonConverter converts arbitrary values from and to JSON (JavaScript Object Notation) strings.
//
// Example:
//
//  value1, _ := convert.FromJson("{\"key\":123}")
//  value2 := convert.JsonConverter.ToMap("{\"key\":123}")
//  value3, _ := convert.ToJson(map[string]int{"key": 123})
//  fmt.Println(value1) // map[key:123]
//  fmt.Println(value2) // map[key:123]
//  fmt.Println(value3) // {"key":123}
var JsonConverter = &_TJsonConverter{}

type _TJsonConverter struct{}

// ToNullableMap converts JSON string into map object or returns null when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or null when conversion is not supported.
func (c *_TJsonConverter) ToNullableMap(value string) *map[string]any {
	return JsonToNullableMap(value)
}

// ToMap converts JSON string into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func (c *_TJsonConverter) ToMap(value string) map[string]any {
	return JsonToMap(value)
}

// ToMapWithDefault converts JSON string into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the JSON string to convert.
//  "defaultValue" - the default value.
// Returns: Map object value or default map when conversion is not supported.
func (c *_TJsonConverter) ToMapWithDefault(value string, defaultValue map[string]any) map[string]any {
	return JsonToMapWithDefault(value, defaultValue)
}

// JsonToNullableMap converts JSON string into map object or returns null when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or null when conversion is not supported.
func JsonToNullableMap(value string) *map[string]any {
	v, _ := FromJson(value)
	if v == nil {
		return nil
	}
	return ToNullableMap(v)
}

// JsonToMap converts JSON string into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func JsonToMap(value string) map[string]any {
	return JsonToMapWithDefault(value, map[string]any{})
}

// JsonToMapWithDefault converts JSON string into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the JSON string to convert.
//  "defaultValue" - the default value.
// Returns: Map object value or default map when conversion is not supported.
func JsonToMapWithDefault(value string, defaultValue map[string]any) map[string]any {
	if m := JsonToNullableMap(value); m != nil {
		return *m
	}
	return defaultValue
}

// FromJson converts value from JSON string
// Parameters: "value" - the JSON string to convert.
// Returns: converted object value or null when value is null.
func FromJson(value string) (any, error) {
	if value == "" {
		return nil, nil
	}

	var m any
	if err := json.Unmarshal([]byte(value), &m); err != nil {
		return nil, err
	}
	return m, nil
}

// ToJson converts value into JSON string.
// Parameters: "value" - the value to convert.
// Returns: JSON string or null when value is null.
func ToJson(value any) (string, error) {
	if value == nil {
		return "", nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}
