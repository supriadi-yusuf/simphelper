package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MappingValue_CollectionIsElementaryType(t *testing.T) {
	_, err := NewCollection(10).MappingValue(10)
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice or map", err.Error())
}

func Test_MappingValue_ParamIsNotFunc(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(10)
	assert.NotNil(t, err)
	assert.Equal(t, "fcriteria argument must be function", err.Error())
}

func Test_MappingValue_FuncHasNoInput(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func() {})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one input parameter", err.Error())
}

func Test_MappingValue_FuncHasOneInput(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(x, y int) {})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one input parameter", err.Error())
}

func Test_MappingValue_FuncHasDifferentInputType(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(checked bool) {})
	assert.NotNil(t, err)
	assert.Equal(t, "element of collection and input parameter for fcriteria are different int type",
		err.Error())
}

func Test_MappingValue_FuncHasNoOutput(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(val int) {})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one output parameter", err.Error())
}

func Test_MappingValue_FuncHasManyOutputs(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).MappingValue(func(val int) (int, string) {
		return 0, "test"
	})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one output parameter", err.Error())
}

func Test_MappingValue_EmptySlice(t *testing.T) {
	result, err := NewCollection([]int{}).MappingValue(func(val int) bool {
		return true
	})
	resultSlice, ok := result.([]bool)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resultSlice))
}

func Test_MappingValue_FromSlice(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3, 4, 5}).MappingValue(func(val int) bool {
		return val%2 == 1
	})
	resultSlice, ok := result.([]bool)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 5, len(resultSlice))
	assert.Equal(t, true, resultSlice[0])
	assert.Equal(t, false, resultSlice[1])
	assert.Equal(t, true, resultSlice[2])
	assert.Equal(t, false, resultSlice[3])
	assert.Equal(t, true, resultSlice[4])
}

func Test_MappingValue_EmptyMap(t *testing.T) {
	result, err := NewCollection(map[string]int{}).MappingValue(func(val int) bool {
		return true
	})
	resultMap, ok := result.(map[string]bool)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resultMap))
}

func Test_MappingValue_FromMap(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}).MappingValue(
		func(val int) bool {
			return val%2 == 1
		})
	resultMap, ok := result.(map[string]bool)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 5, len(resultMap))
	assert.Equal(t, true, resultMap["1"])
	assert.Equal(t, false, resultMap["2"])
	assert.Equal(t, true, resultMap["3"])
	assert.Equal(t, false, resultMap["4"])
	assert.Equal(t, true, resultMap["5"])
}
