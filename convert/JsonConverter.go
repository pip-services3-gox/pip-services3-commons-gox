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
var JsonConverter = &_TJsonConverter{
	_jsonEngine: &defaultJsonEngine{},
}

type _TJsonConverter struct {
	_jsonEngine IJSONEngine
}

// IJSONEngine interface which helps to marshal and unmarshal json
type IJSONEngine interface {
	FromJson(value string) (any, error)
	ToJson(value any) (string, error)
}

// SetUpJSONEngine allows overriding JSON engine
func (j *_TJsonConverter) SetUpJSONEngine(jsonEngine IJSONEngine) bool {
	if jsonEngine == nil {
		return false
	}
	j._jsonEngine = jsonEngine
	return true
}

// FromJson converts value from JSON string
// Parameters: "value" - the JSON string to convert.
// Returns: converted object value or null when value is null.
func (j *_TJsonConverter) FromJson(value string) (any, error) {
	return j._jsonEngine.FromJson(value)
}

// ToJson converts value into JSON string.
// Parameters: "value" - the value to convert.
// Returns: JSON string or null when value is null.
func (j *_TJsonConverter) ToJson(value any) (string, error) {
	return j._jsonEngine.ToJson(value)
}

// ToNullableMap converts JSON string into map object or returns null when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value and true or null and false when conversion is not supported.
func (j *_TJsonConverter) ToNullableMap(value string) (map[string]any, bool) {
	v, _ := j.FromJson(value)
	if v == nil {
		return nil, false
	}
	return toNullableMap(v)
}

// ToMap converts JSON string into map object or returns empty map when conversion is not possible.
// Parameters: "value" - the JSON string to convert.
// Returns: Map object value or empty map when conversion is not supported.
func (j *_TJsonConverter) ToMap(value string) map[string]any {
	return j.ToMapWithDefault(value, map[string]any{})
}

// ToMapWithDefault converts JSON string into map object or returns default map when conversion is not possible.
// Parameters:
//  "value" - the JSON string to convert.
//  "defaultValue" - the default value.
// Returns: Map object value or default map when conversion is not supported.
func (j *_TJsonConverter) ToMapWithDefault(value string, defaultValue map[string]any) map[string]any {
	if m, ok := j.ToNullableMap(value); ok {
		return m
	}
	return defaultValue
}

// jsonEngine is a default json converter engine
type defaultJsonEngine struct{}

// FromJson converts value from JSON string
// Parameters: "value" - the JSON string to convert.
// Returns: converted object value or null when value is null.
func (je defaultJsonEngine) FromJson(value string) (any, error) {
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
func (je defaultJsonEngine) ToJson(value any) (string, error) {
	if value == nil {
		return "", nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}
