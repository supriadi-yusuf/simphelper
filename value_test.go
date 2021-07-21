package simphelper

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsIn_String(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	_, err := newValue.IsIn("test is in")
	assert.NotNil(t, err, "it should be error")
}

func Test_IsIn_BoolSlice(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	_, err := newValue.IsIn([]bool{false, true, true, false})
	assert.NotNil(t, err, "it should be error")
}

func Test_IsIn_EmptyIntMap(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	result, err := newValue.IsIn(map[string]int{})
	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, result, false, "it should be false")
}

func Test_IsIn_NotExistInIntMap(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	result, err := newValue.IsIn(map[string]int{"satu": 1, "dua": 2, "tiga": 3})
	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, result, false, "it should be false")
}

func Test_IsIn_ExistInIntMap(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	result, err := newValue.IsIn(map[string]int{"1": 1, "2": 2, "3": 3, "10": 10})
	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, result, true, "it should be true")
}

func Test_IsIn_EmptyIntSlice(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	result, err := newValue.IsIn([]int{})
	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, result, false, "it should be false")
}

func Test_IsIn_NotExistInIntSlice(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	result, err := newValue.IsIn([]int{1, 2, 3})
	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, result, false, "it should be false")
}

func Test_IsIn_ExistInIntSlice(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	result, err := newValue.IsIn([]int{1, 2, 3, 10, 11})
	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, result, true, "it should be true")
}

func Test_IsNumber_int(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_int8(t *testing.T) {
	var i int8 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_int16(t *testing.T) {
	var i int16 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_int32(t *testing.T) {
	var i int32 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_int64(t *testing.T) {
	var i int64 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_uint(t *testing.T) {
	var i uint = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_uint8(t *testing.T) {
	var i uint8 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_uint16(t *testing.T) {
	var i uint16 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_uint32(t *testing.T) {
	var i uint32 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_uint64(t *testing.T) {
	var i uint64 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_float32(t *testing.T) {
	var i float32 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_float64(t *testing.T) {
	var i float64 = 10

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, true, "it should be true")
}

func Test_IsNumber_string(t *testing.T) {
	var i string = "10"

	newValue := NewValue(i)
	isNumber := newValue.IsNumber()

	assert.Equal(t, isNumber, false, "it should be false")
}

func Test_ToFloat_int8(t *testing.T) {
	var i int8 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_int16(t *testing.T) {
	var i int16 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_int32(t *testing.T) {
	var i int32 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_int64(t *testing.T) {
	var i int64 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_int(t *testing.T) {
	var i int = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_uint(t *testing.T) {
	var i uint = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_uint8(t *testing.T) {
	var i uint8 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_uint16(t *testing.T) {
	var i uint16 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_uint32(t *testing.T) {
	var i uint32 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_uint64(t *testing.T) {
	var i uint64 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_float32(t *testing.T) {
	var i float32 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_float64(t *testing.T) {
	var i float64 = 10

	newValue := NewValue(i)
	f, err := newValue.ToFloat()

	assert.Nil(t, err, "it should be not error")
	assert.Equal(t, f, float64(10), "f should be 10")
	assert.Equal(t, reflect.TypeOf(f).Name(), "float64", "f type should be float64")
}

func Test_ToFloat_string(t *testing.T) {
	var i string = "10"

	newValue := NewValue(i)
	_, err := newValue.ToFloat()

	assert.NotNil(t, err, "it should be error")
}

func Test_NewValue(t *testing.T) {
	newValue := NewValue(10)
	assert.NotNil(t, newValue, "it should be not nil")
}
