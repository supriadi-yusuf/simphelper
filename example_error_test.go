package simphelper

import (
	"fmt"
)

func ExampleGetErrorOnPanic() {

	result := func() (err error) {

		defer GetErrorOnPanic(&err)

		fmt.Println("Hello")

		return
	}()
	if result != nil {
		fmt.Println(result.Error())
	}
}
