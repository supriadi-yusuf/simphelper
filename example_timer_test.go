package simphelper

import (
	"time"
)

func ExampleWait() {

	Wait(5 * time.Second) // sleep / wait for 5 seconds
}
