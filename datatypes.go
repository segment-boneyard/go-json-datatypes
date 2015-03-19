package datatypes

import "reflect"

// FYI these are the types when you use json.Unmarshal
//
// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null

var byteSliceType = reflect.TypeOf([]byte(nil))

func DataType(v interface{}) string {
	if v == nil {
		return "null"
	}

	valueOf := reflect.ValueOf(v)

	switch valueOf.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "number"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "number"
	case reflect.Float32, reflect.Float64:
		return "number"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "boolean"
	case reflect.Array, reflect.Slice:
		if valueOf.Type() == byteSliceType {
			return "string"
		}
		return "array"
	case reflect.Map, reflect.Struct:
		return "object"
	default:
		return "value"
	}
}
