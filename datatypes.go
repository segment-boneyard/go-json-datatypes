package datatypes

import "reflect"

var typeMap map[string]string = map[string]string{
	"bool": "Boolean",

	"float64": "Number",
	"int":     "Number",

	"string": "String",

	"[]interface{}": "Array",
	"[]string":      "Array",

	"map[string]interface {}": "Object",
	"nil": "null",
}

func DataType(obj interface{}) string {
	if obj == nil {
		return "null"
	}

	t := reflect.TypeOf(obj).String()

	return typeMap[t]
}
