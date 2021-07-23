package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//NewCollection is function creating object whose type is ICollection
//
//data must be array, slice, map or struct
func NewCollectionx(data interface{}) ICollection {
	collection := new(myCollection)
	collection.data = data
	return collection
}

func Test_Collection(t *testing.T) {
	col := NewCollection(10)
	assert.NotNil(t, col, "it should be not nil")
}
