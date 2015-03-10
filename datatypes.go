package datatypes

func DataType(v interface{}) string {
	if v == nil {
		return "null"
	}

	switch v.(type) {
	case int, int8, int32, int64, uint, uint8, uint32, uint64, float32, float64:
		return "Number"
	case string:
		return "String"
	case bool:
		return "Boolean"
	case []interface{}, []string, []int, []int8, []int32, []int64, []uint, []uint8, []uint32, []uint64, []float32, []float64:
		return "Array"
	case map[interface{}]interface{}, map[string]interface{}:
		return "Object"
	default:
		return "Value"
	}
}
