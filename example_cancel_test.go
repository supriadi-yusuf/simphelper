package simphelper

import (
	"context"
	"time"
)

func ExampleCheckForCancelation() {

	var maxCntr = 10000
	var cntr int = 0

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // do cancelation

	go func() {

		defer DoCancelRecover()

		for cntr = 0; cntr < maxCntr; cntr++ {
			Wait(10 * time.Millisecond)
			CheckForCancelation(ctx)
		}
	}()

	Wait(1000 * time.Millisecond)

}
