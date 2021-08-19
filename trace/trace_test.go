package trace

import (
	"time"
)

func ExampleExecuteTime() {
	ExecuteTime("step 1.", func() {
		time.Sleep(200 * time.Millisecond)
	})

	ExecuteTime("step 2.", func() {
		time.Sleep(300 * time.Millisecond)
	})

	ExecuteTime("step 3.", func() {
		time.Sleep(400 * time.Millisecond)
	})

	// Output:
}
