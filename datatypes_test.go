package datatypes

import "github.com/bmizerany/assert"
import "testing"

func TestDataTypeForArrays(t *testing.T) {
	assert.Equal(t, "array", DataType([]bool{true, false}))
	assert.Equal(t, "array", DataType([2]int{1, 2}))
	assert.Equal(t, "array", DataType([]string{"hello", "world"}))

	// Byte slices are base64 encoded; byte arrays are not.
	assert.Equal(t, "string", DataType([]byte{1, 0, 1}))
	assert.Equal(t, "string", DataType([]uint8{1, 0, 1}))
	assert.Equal(t, "array", DataType([3]byte{1, 0, 1}))
}

type sample struct {
	id   int
	name string
}

func TestDataTypeForObjects(t *testing.T) {
	assert.Equal(t, "object", DataType(sample{id: 1, name: "test"}))
	assert.Equal(t, "object", DataType(map[string]interface{}{"hello": "world"}))
}

func TestDataTypeForPrimitives(t *testing.T) {
	assert.Equal(t, "boolean", DataType(true))
	assert.Equal(t, "number", DataType(16))
	assert.Equal(t, "number", DataType(1.1))
	assert.Equal(t, "number", DataType(-1))
	assert.Equal(t, "string", DataType("hello world"))
	assert.Equal(t, "null", DataType(nil))
}
