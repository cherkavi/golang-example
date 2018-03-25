package timer

import (
	"fmt"
	"time"
)

func ExampleTicker() {
	ticker := time.NewTicker(time.Millisecond * 200)

	for stopAfterCount := 5; stopAfterCount > 0; stopAfterCount-- {
		tick := <-ticker.C
		fmt.Printf(" next tick done: %v \n", tick)
		if stopAfterCount == 3 {
			fmt.Println(" stop ticker ")
			ticker.Stop()
			break
		}
	}
}
