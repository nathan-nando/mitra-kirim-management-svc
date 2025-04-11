package converter

import (
	"fmt"
	"reflect"
	"strings"
)

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func ConvertKeyValue(obj interface{}) []KeyValue {
	v := reflect.ValueOf(obj)

	// Handle pointer input by dereferencing
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	t := v.Type()

	// Check if the type is a struct
	if t.Kind() != reflect.Struct {
		return nil
	}

	var result []KeyValue
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")

		// Ignore fields without JSON tags or with "-"
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Extract just the field name from the json tag (remove options like omitempty)
		jsonFieldName := strings.Split(jsonTag, ",")[0]

		// Skip if the json field name is empty (another way to specify ignore)
		if jsonFieldName == "" {
			continue
		}

		// Get field value
		value := v.Field(i).Interface()

		// Handle nil pointers gracefully
		if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
			value = nil
		}

		// Skip empty strings
		if strValue, ok := value.(string); ok && strValue == "" {
			continue
		}

		// Skip nil pointers to strings
		if strPtr, ok := value.(*string); ok && (strPtr == nil || *strPtr == "") {
			continue
		}

		result = append(result, KeyValue{
			Key:   jsonFieldName,
			Value: value,
		})
	}

	return result
}

func ConvertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case *string:
		if v != nil {
			return *v
		}
		return ""
	case int, int64, float64, bool:
		return fmt.Sprintf("%v", v) // Convert numbers, bools to string
	default:
		return fmt.Sprintf("%v", v) // Fallback for unknown types
	}
}
