package datatypes

// FYI these are the types when you use json.Unmarshal
//
// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null

func DataType(v interface{}) string {
	if v == nil {
		return "null"
	}

	switch v.(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64, float32, float64:
		return "number"
	case string:
		return "string"
	case bool:
		return "boolean"
	case []interface{}, []string, []int, []int8, []int32, []int64, []uint, []uint8, []uint32, []uint64, []float32, []float64:
		return "array"
	case map[interface{}]interface{}, map[string]interface{}:
		return "object"
	default:
		return "value"
	}
}
