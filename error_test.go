package simphelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetErrorOnPanic_error(t *testing.T) {

	conversion := func(x interface{}) int {
		return x.(int)
	}

	err := func() (err error) {
		defer GetErrorOnPanic(&err)
		var val float32 = 10
		conversion(val)
		return
	}()

	assert.NotNil(t, err, "it should be error")
}

func Test_GetErrorOnPanic_noerror(t *testing.T) {

	conversion := func(x interface{}) int {
		return x.(int)
	}

	err := func() (err error) {
		defer GetErrorOnPanic(&err)
		var val int = 10
		conversion(val)
		return
	}()

	assert.Nil(t, err, "it should be error")
}
