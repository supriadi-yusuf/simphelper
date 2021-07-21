package simphelper

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// CheckForCancelation is function to check if current process is being canceled outside
func Test_CheckForCancelation_cancel(t *testing.T) {
	err := func() (err error) {
		defer func() {
			if it := recover(); it != nil {
				err = fmt.Errorf("%v", it)
			}
		}()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		CheckForCancelation(ctx)
		return
	}()

	assert.NotNil(t, err, "it should be error")
}

func Test_CheckForCancelation_nocancel(t *testing.T) {
	err := func() (err error) {
		defer func() {
			if it := recover(); it != nil {
				err = fmt.Errorf("%v", it)
			}
		}()

		CheckForCancelation(context.Background())
		return
	}()

	assert.Nil(t, err, "it should be error")
}

// DoCancelRecover is function to stop panic propagation
func Test_DoCancelRecover_error(t *testing.T) {

	testConversion := func() (err error) {

		defer func() {
			if it := recover(); it != nil {
				err = fmt.Errorf("%v", it)
			}
		}()

		conversion := func(x interface{}) int {
			defer DoCancelRecover()

			return x.(int)
		}

		var val float64 = 10
		conversion(val)

		return
	}

	err := testConversion()
	assert.Nil(t, err, "it shoul be no error")
}

func Test_DoCancelRecover_noerror(t *testing.T) {

	testConversion := func() (err error) {

		defer func() {
			if it := recover(); it != nil {
				err = fmt.Errorf("%v", it)
			}
		}()

		conversion := func(x interface{}) int {
			defer DoCancelRecover()

			return x.(int)
		}

		var val int = 10
		conversion(val)

		return
	}

	err := testConversion()
	assert.Nil(t, err, "it shoul be no error")
}
