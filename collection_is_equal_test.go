package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collection_IsEqual_BothCollectionAndParamAreElementaryType(t *testing.T) {

	_, err := NewCollection(10).IsEqual(10)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "collection must be array, slice, struct or map")
}

func Test_Collection_IsEqual_ParamAreElementaryType(t *testing.T) {

	_, err := NewCollection([]int{1, 2, 3, 10}).IsEqual(10)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "types are different")
}

func Test_Collection_IsEqual_CollectionIsSliceAndParamIsMap(t *testing.T) {

	_, err := NewCollection([]int{1, 2, 3, 10}).IsEqual(map[string]int{"1": 1, "2": 2, "3": 3, "10": 10})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "types are different")
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
	assert.Equal(t, err.Error(), "types are different")
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
	assert.Equal(t, err.Error(), "types are different")
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
	assert.Equal(t, err.Error(), "types are different")
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
	assert.Equal(t, result, false)
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
	assert.Equal(t, result, true)
}
