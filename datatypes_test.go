package datatypes

import "github.com/bmizerany/assert"
import "testing"

func TestDatatype(t *testing.T) {
	assert.Equal(t, "Boolean", DataType(true))
	assert.Equal(t, "Number", DataType(16))
	assert.Equal(t, "String", DataType("hello world"))
	assert.Equal(t, "Array", DataType([]string{"hello", "world"}))
	assert.Equal(t, "Object", DataType(map[string]interface{}{"hello": "world"}))
	assert.Equal(t, "null", DataType(nil))
}
