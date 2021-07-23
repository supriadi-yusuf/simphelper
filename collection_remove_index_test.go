package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collection_RemoveIndex_CollectionisElementaryType(t *testing.T) {
	_, err := NewCollection(10).RemoveIndex(10)
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice or map", err.Error())
}

func Test_Collection_RemoveIndex_IndexIsNotElementaryType(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3}).RemoveIndex([]int{1, 2, 3})
	assert.NotNil(t, err)
	assert.Equal(t, "index must be not array, slice, map or struct", err.Error())
}

func Test_Collection_RemoveIndex_EmptyMap(t *testing.T) {
	result, err := NewCollection(map[int]string{}).RemoveIndex(1)
	mapResult, ok := result.(map[int]string)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(mapResult))
}

func Test_Collection_RemoveIndex_MapKeyIsDifferent(t *testing.T) {
	_, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).RemoveIndex(1)
	assert.NotNil(t, err)
	assert.Equal(t, "map key and index have different data type", err.Error())
}

func Test_Collection_RemoveIndex_IndexIsNotInMap(t *testing.T) {
	result, err := NewCollection(map[int]string{1: "1", 2: "2", 3: "3"}).RemoveIndex(10)
	mapResult, ok := result.(map[int]string)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, len(mapResult))
}

func Test_Collection_RemoveIndex_IndexIsInMap(t *testing.T) {
	result, err := NewCollection(map[int]string{1: "1", 2: "2", 3: "3"}).RemoveIndex(2)
	mapResult, ok := result.(map[int]string)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, len(mapResult))
}

func Test_Collection_RemoveIndex_EmptyArray(t *testing.T) {
	result, err := NewCollection([]string{}).RemoveIndex(1)
	mapResult, ok := result.([]string)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(mapResult))
}

func Test_Collection_RemoveIndex_ArrayKeyIsDifferenTypetWithIndex(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3}).RemoveIndex("1")
	assert.NotNil(t, err)
	assert.Equal(t, "interface conversion: interface {} is string, not int", err.Error())
}

func Test_Collection_RemoveIndex_IndexIsNotInSlice(t *testing.T) {
	result, err := NewCollection([]string{"1", "2", "3"}).RemoveIndex(10)
	mapResult, ok := result.([]string)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, len(mapResult))
}

func Test_Collection_RemoveIndex_IndexIsInSlice(t *testing.T) {
	result, err := NewCollection([]string{"1", "2", "3"}).RemoveIndex(2)
	mapResult, ok := result.([]string)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, len(mapResult))
}
