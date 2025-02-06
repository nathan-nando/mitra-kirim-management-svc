package converter

import (
	"fmt"
	"reflect"
)

type KeyValue struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func ConvertKeyValue(obj interface{}) []KeyValue {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

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

		value := v.Field(i).Interface() // Get field value

		// Handle nil pointers gracefully
		if v.Field(i).Kind() == reflect.Ptr && v.Field(i).IsNil() {
			value = nil
		}

		result = append(result, KeyValue{
			Key:   jsonTag,
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
