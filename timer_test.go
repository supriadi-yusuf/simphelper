package simphelper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Wait(t *testing.T) {

	d1 := 5 * time.Second
	startTime := time.Now()
	Wait(d1)
	d2 := time.Since(startTime)

	fsecond2 := d2.Seconds()

	//dseconds := fsecond2 - fsecond1
	assert.GreaterOrEqual(t, fsecond2, float64(5), "delay time should not be less than 5 seconds")
}
