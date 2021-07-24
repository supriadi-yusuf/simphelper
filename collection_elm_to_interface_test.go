package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvElmToInterface_CollectionIsElementaryType(t *testing.T) {
	_, err := NewCollection(10).ConvElmToInterface()
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice or map", err.Error())
}

func Test_ConvElmToInterface_EmptyMap(t *testing.T) {
	result, err := NewCollection(map[string]int{}).ConvElmToInterface()
	resultMap, ok := result.(map[string]interface{})
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resultMap))
}

func Test_ConvElmToInterface_FromMap(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}).ConvElmToInterface()
	resultMap, ok := result.(map[string]interface{})
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 5, len(resultMap))
}

func Test_ConvElmToInterface_EmptySlice(t *testing.T) {
	result, err := NewCollection([]int{}).ConvElmToInterface()
	resultSlice, ok := result.([]interface{})
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resultSlice))
}

func Test_ConvElmToInterface_FromSlice(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3, 4, 5}).ConvElmToInterface()
	resultSlice, ok := result.([]interface{})
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 5, len(resultSlice))
}
