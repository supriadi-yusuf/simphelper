package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collection_IsEqual_BothCollectionAndParamAreElementaryType(t *testing.T) {

	_, err := NewCollection(10).IsEqual(10)
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice, struct or map", err.Error())
}

func Test_Collection_IsEqual_ParamAreElementaryType(t *testing.T) {

	_, err := NewCollection([]int{1, 2, 3, 10}).IsEqual(10)
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_Collection_IsEqual_CollectionIsSliceAndParamIsMap(t *testing.T) {

	_, err := NewCollection([]int{1, 2, 3, 10}).IsEqual(map[string]int{"1": 1, "2": 2, "3": 3, "10": 10})
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_Collection_IsEqual_StructIsDifferentInFieldNumber(t *testing.T) {

	_, err := NewCollection(
		struct {
			Name string
			Age  int
		}{"supriadi", 40}).
		IsEqual(
			struct {
				Name    string
				Age     int
				citizen string
			}{"supriadi", 40, "indonesia"},
		)
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_Collection_IsEqual_StructIsDifferentInField(t *testing.T) {

	_, err := NewCollection(
		struct {
			Name string
			Age  int
		}{"supriadi", 40}).
		IsEqual(
			struct {
				Name   string
				Weight int
			}{"supriadi", 40},
		)
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_Collection_IsEqual_StructIsDifferentInFieldSequence(t *testing.T) {

	_, err := NewCollection(
		struct {
			Name string
			Age  int
		}{"supriadi", 40}).
		IsEqual(
			struct {
				Age  int
				Name string
			}{40, "supriadi"},
		)
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_Collection_IsEqual_StructIsDifferentInValues(t *testing.T) {

	result, err := NewCollection(
		struct {
			Name string
			Age  int
		}{"supriadi", 40}).
		IsEqual(
			struct {
				Name string
				Age  int
			}{"supriadi", 41},
		)
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_StructIsSame(t *testing.T) {

	result, err := NewCollection(
		struct {
			Name string
			Age  int
		}{"supriadi", 40}).
		IsEqual(
			struct {
				Name string
				Age  int
			}{"supriadi", 40},
		)
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_Collection_IsEqual_SliceWithDifferentLength(t *testing.T) {

	result, err := NewCollection([]int{1, 2, 3}).IsEqual([]int{1, 2, 3, 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_SliceWithDifferentContent(t *testing.T) {

	result, err := NewCollection([]int{1, 2, 3}).IsEqual([]int{1, 2, 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_SliceWithSameContent(t *testing.T) {

	result, err := NewCollection([]int{1, 2, 3}).IsEqual([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_Collection_IsEqual_SliceWithSameContentButDifferentSequence(t *testing.T) {

	result, err := NewCollection([]int{1, 2, 3}).IsEqual([]int{1, 3, 2})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_MapWithDifferentLength(t *testing.T) {

	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).
		IsEqual(map[string]int{"1": 1, "2": 2, "3": 3, "4": 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_MapWithDifferentIndex(t *testing.T) {

	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).
		IsEqual(map[string]int{"1": 1, "2": 2, "4": 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_MapWithDifferentContent(t *testing.T) {

	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).
		IsEqual(map[string]int{"1": 1, "2": 2, "3": 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_Collection_IsEqual_MapWithSameContent(t *testing.T) {

	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).
		IsEqual(map[string]int{"1": 1, "2": 2, "3": 3})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_Collection_IsEqual_MapWithSameContentButDifferentSequence(t *testing.T) {

	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).
		IsEqual(map[string]int{"1": 1, "3": 3, "2": 2})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}
