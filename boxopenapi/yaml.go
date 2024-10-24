package boxopenapi

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
// Función recursiva para convertir un objeto JSON a una cadena YAML
func convertToYAML1(obj interface{}, indent int) string {
	indentStr := strings.Repeat("  ", indent)
	yamlStr := ""

	switch v := obj.(type) {
	case map[string]interface{}:
		sorted := make([]string, 0, len(v))
		for key := range v {
			sorted = append(sorted, key)
		}
		sort.Strings(sorted)

		for _, key := range sorted {
			value := v[key]

			if isContainer(value) {
				yamlStr += fmt.Sprintf("%s%s:\n%s", indentStr, key, convertToYAML(value, indent+1))
			} else {
				yamlStr += fmt.Sprintf("%s%s: %s\n", indentStr, key, convertToYAML(value, indent+1))
			}
		}
	case []interface{}:
		for _, value := range v {
			if isContainer(value) {
				yamlStr += fmt.Sprintf("%s- %s", indentStr, convertToYAML(value, indent+1))
			} else {
				yamlStr += fmt.Sprintf("%s- %s\n", indentStr, convertToYAML(value, indent+1))
			}
		}
	// case []map[string]interface{}:
	// 	for _, value := range v {
	// 		if isContainer(value) {
	// 			yamlStr += fmt.Sprintf("%s- %s", indentStr, convertToYAML(value, indent+1))
	// 		} else {
	// 			yamlStr += fmt.Sprintf("%s- %s\n", indentStr, convertToYAML(value, indent+1))
	// 		}
	// 	}
	case string:
		yamlStr += fmt.Sprintf("%s", v)
		// yamlStr += fmt.Sprintf("%s%s\n", indentStr, v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		// yamlStr += fmt.Sprintf("%v", v)
		yamlStr += fmt.Sprintf("%v", v)
	case bool:
		// yamlStr += fmt.Sprintf("%v", v)
		yamlStr += fmt.Sprintf("%v", v)
	default:
		// yamlStr += "null"
		yamlStr += fmt.Sprintf("null")
	}

	return yamlStr
}
*/

func isContainer(v any) bool {
	switch v.(type) {
	case []map[string]interface{}:
		return true
	case map[string]interface{}:
		return true
	case []interface{}:
		return true
	default:
		return false
	}
}

// ConvertToYAML convierte un objeto Go a su representación YAML.
func ConvertToYAML(obj interface{}) string {
	return convertToYAML(reflect.ValueOf(obj), 0)
}

func convertToYAML(value reflect.Value, indent int) string {
	var result strings.Builder
	indentation := strings.Repeat("  ", indent)

	for value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.Map:

		keys := value.MapKeys()
		sortedKeys := make([]string, 0, len(keys))
		keysPosition := map[string]reflect.Value{}
		for _, key := range keys {
			s := fmt.Sprintf("%v", key.Interface())
			sortedKeys = append(sortedKeys, s)
			keysPosition[s] = key
		}
		sort.Strings(sortedKeys)

		for _, key := range sortedKeys {
			result.WriteString(fmt.Sprintf("%s%s: %s\n", indentation, key, convertToYAML(value.MapIndex(keysPosition[key]), indent+1)))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fieldName := value.Type().Field(i).Name
			result.WriteString(fmt.Sprintf("%s%s: %s\n", indentation, fieldName, convertToYAML(value.Field(i), indent+1)))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			result.WriteString(fmt.Sprintf("%s- %s\n", indentation, convertToYAML(value.Index(i), indent+1)))
		}
	case reflect.String:
		result.WriteString(fmt.Sprintf("%s", value.String()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result.WriteString(fmt.Sprintf("%d", value.Int()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result.WriteString(fmt.Sprintf("%d", value.Uint()))
	case reflect.Float32, reflect.Float64:
		result.WriteString(fmt.Sprintf("%f", value.Float()))
	case reflect.Bool:
		result.WriteString(fmt.Sprintf("%t", value.Bool()))
	default:
		result.WriteString(fmt.Sprintf("%v", value.Interface()))
	}

	return result.String()
}
