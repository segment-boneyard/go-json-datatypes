package datatypes

import "github.com/bmizerany/assert"
import "testing"

func TestDatatype(t *testing.T) {
	assert.Equal(t, "boolean", DataType(true))
	assert.Equal(t, "number", DataType(16))
	assert.Equal(t, "string", DataType("hello world"))
	assert.Equal(t, "array", DataType([]string{"hello", "world"}))
	assert.Equal(t, "object", DataType(map[string]interface{}{"hello": "world"}))
	assert.Equal(t, "null", DataType(nil))
}
