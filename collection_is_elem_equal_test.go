package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsElemEqual_CollectionIsElementaryType(t *testing.T) {
	_, err := NewCollection(10).IsElemEqual(0)
	assert.NotNil(t, err)
	assert.Equal(t, "collection must be array, slice, struct or map", err.Error())
}

func Test_IsElemEqual_CollectionAndParamIsDifferent(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3}).IsElemEqual(0)
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_IsElemEqual_StructUnExportedElement(t *testing.T) {
	_, err := NewCollection(
		struct {
			Name   string
			Age    int
			weight int
		}{"supri", 40, 65}).IsElemEqual(
		struct {
			Name   string
			Age    int
			weight int
		}{"supri", 40, 65})
	assert.NotNil(t, err)
	assert.Equal(t, "reflect.Value.Interface: cannot return value obtained from unexported field or method",
		err.Error())
}

func Test_IsElemEqual_DifferentStructElementSequence(t *testing.T) {
	_, err := NewCollection(
		struct {
			Name   string
			Age    int
			Weight int
		}{"supri", 40, 65}).IsElemEqual(
		struct {
			Name   string
			Weight int
			Age    int
		}{"supri", 65, 40})
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_IsElemEqual_DifferentStructElementValue(t *testing.T) {
	result, err := NewCollection(
		struct {
			Name   string
			Age    int
			Weight int
		}{"supri", 40, 65}).IsElemEqual(
		struct {
			Name   string
			Age    int
			Weight int
		}{"supri", 41, 65})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_SameStructElementValue(t *testing.T) {
	result, err := NewCollection(
		struct {
			Name   string
			Age    int
			Weight int
		}{"supri", 40, 65}).IsElemEqual(
		struct {
			Name   string
			Age    int
			Weight int
		}{"supri", 40, 65})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_SliceAndElementaryType(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3}).IsElemEqual(0)
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_IsElemEqual_SliceAndStruct(t *testing.T) {
	_, err := NewCollection([]int{1, 2, 3}).IsElemEqual(struct{}{})
	assert.NotNil(t, err)
	assert.Equal(t, "types are different", err.Error())
}

func Test_IsElemEqual_SliceWithDifferentSize(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3}).IsElemEqual([]int{1, 2, 3, 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_SliceWithCollectionNotInParam(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3}).IsElemEqual([]int{1, 2, 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_SliceWithParamNotInCollectionParam(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 2}).IsElemEqual([]int{1, 2, 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_SliceWithSameEmptyValue(t *testing.T) {
	result, err := NewCollection([]int{}).IsElemEqual([]int{})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_SliceWithSameValue(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3}).IsElemEqual([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_SliceWithSameValueButDifferentSequence(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 3}).IsElemEqual([]int{1, 3, 2})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_SliceWithDifferentValue(t *testing.T) {
	result, err := NewCollection([]int{1, 2, 2, 3}).IsElemEqual([]int{1, 2, 3, 3})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_MapWithDifferentSize(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).IsElemEqual(
		map[string]int{"1": 1, "2": 2, "3": 3, "4": 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_MapWithCollectionNotInParam(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).IsElemEqual(
		map[string]int{"1": 1, "2": 2, "4": 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_MapWithParamNotInCollectionParam(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 2}).IsElemEqual(
		map[string]int{"1": 1, "2": 2, "3": 4})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}

func Test_IsElemEqual_MapWithSameEmptyValue(t *testing.T) {
	result, err := NewCollection(map[string]int{}).IsElemEqual(map[string]int{})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_MapWithSameValue(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).IsElemEqual(
		map[string]int{"1": 1, "2": 2, "3": 3})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_MapWithSameValueButDifferentSequence(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).IsElemEqual(
		map[string]int{"1": 1, "3": 3, "2": 2})
	assert.Nil(t, err)
	assert.Equal(t, true, result)
}

func Test_IsElemEqual_MapWithDifferentValue(t *testing.T) {
	result, err := NewCollection(map[string]int{"1": 1, "2": 2, "3": 3}).IsElemEqual(
		map[string]int{"1": 1, "2": 3, "3": 2})
	assert.Nil(t, err)
	assert.Equal(t, false, result)
}
