package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MeanValue_CollectionIsElementaryType(t *testing.T) {
	_, err := NewCollection(10).MeanValue()
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice or map", err.Error())
}

func Test_MeanValue_CollectionIsEmptySlice(t *testing.T) {
	result, err := NewCollection([]int{}).MeanValue()
	assert.Nil(t, err)
	assert.Equal(t, float64(0), result)
}

func Test_MeanValue_CollectionIsEmptySliceAndNotNumber(t *testing.T) {
	_, err := NewCollection([]bool{}).MeanValue()
	assert.NotNil(t, err)
	assert.Equal(t, "element type must be number", err.Error())
}

func Test_MeanValue_SliceElementIsNotNumber(t *testing.T) {
	_, err := NewCollection([]bool{true, false, true, false, true}).MeanValue()
	assert.NotNil(t, err)
	assert.Equal(t, "element type must be number", err.Error())
}

func Test_MeanValue_SliceElementIsNumber(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3, 4, 5}).MeanValue()

	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)
}

func Test_MeanValue_EmptyMap(t *testing.T) {
	result, err := NewCollection(map[int]int{}).MeanValue()
	assert.Nil(t, err)
	assert.Equal(t, 0.0, result)
}

func Test_MeanValue_EmptyMapAndNotNumber(t *testing.T) {
	_, err := NewCollection(map[int]bool{}).MeanValue()
	assert.NotNil(t, err)
	assert.Equal(t, "element type must be number", err.Error())
}

func Test_MeanValue_MapElementIsNotNumber(t *testing.T) {
	_, err := NewCollection(map[int]bool{1: true, 2: false, 3: true, 4: false, 5: true}).MeanValue()
	assert.NotNil(t, err)
	assert.Equal(t, "element type must be number", err.Error())
}

func Test_MeanValue_MapElementIsNumber(t *testing.T) {
	result, err := NewCollection(map[string]uint{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}).MeanValue()

	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)
}
