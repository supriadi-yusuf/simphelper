package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FilterValue_CollectionIsElementaryType(t *testing.T) {
	_, err := NewCollection(10).FilterValue(10)
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice or map", err.Error())
}

func Test_FilterValue_ParamIsNotFunc(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(10)
	assert.NotNil(t, err)
	assert.Equal(t, "fcriteria argument must be function", err.Error())
}

func Test_FilterValue_FuncHasNoInput(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func() {})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one input parameter", err.Error())
}

func Test_FilterValue_FuncHasOneInput(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func(x, y int) {})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one input parameter", err.Error())
}

func Test_FilterValue_FuncHasDifferentInputType(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func(checked bool) {})
	assert.NotNil(t, err)
	assert.Equal(t, "element of collection and input parameter for fcriteria are different int type",
		err.Error())
}

func Test_FilterValue_FuncHasNoOutput(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func(val int) {})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one output parameter", err.Error())
}

func Test_FilterValue_FuncHasManyOutputs(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func(val int) (int, string) {
		return 0, "test"
	})
	assert.NotNil(t, err)
	assert.Equal(t, "fcritera must have only one output parameter", err.Error())
}

func Test_FilterValue_FuncTypeIsNotBoolean(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3, 4, 5, 6}).FilterValue(func(val int) int {
		return 0
	})
	assert.NotNil(t, err)
	assert.Equal(t, "output parameter for fcriteria must be boolean", err.Error())
}

func Test_FilterValue_SliceIsEmpty(t *testing.T) {
	result, err := NewCollection([]int{}).FilterValue(func(val int) bool {
		return true
	})
	resulSlice, ok := result.([]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resulSlice))
}

func Test_FilterValue_InSlice_CriteriaIsTrue(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3, 4, 5}).FilterValue(func(val int) bool {
		return true
	})
	resulSlice, ok := result.([]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 5, len(resulSlice))
}

func Test_FilterValue_InSlice_CriteriaIsFalse(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3, 4, 5}).FilterValue(func(val int) bool {
		return false
	})
	resulSlice, ok := result.([]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resulSlice))
}

func Test_FilterValue_InSlice(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3, 4, 5}).FilterValue(func(val int) bool {
		return val%2 == 1
	})
	resulSlice, ok := result.([]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, len(resulSlice))
}

func Test_FilterValue_MapIsEmpty(t *testing.T) {
	result, err := NewCollection(map[string]int{}).FilterValue(func(val int) bool {
		return true
	})
	resulSlice, ok := result.(map[string]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resulSlice))
}

func Test_FilterValue_InMap_CriteriaIsTrue(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}).FilterValue(func(val int) bool {
		return true
	})
	resulSlice, ok := result.(map[string]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 5, len(resulSlice))
}

func Test_FilterValue_InMap_CriteriaIsFalse(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}).FilterValue(
		func(val int) bool {
			return false
		})
	resulSlice, ok := result.(map[string]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(resulSlice))
}

func Test_FilterValue_InMap(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5}).FilterValue(
		func(val int) bool {
			return val%2 == 1
		})
	resulSlice, ok := result.(map[string]int)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, len(resulSlice))
}
