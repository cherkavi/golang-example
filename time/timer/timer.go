package timer

import (
	"fmt"
	"time"
)

func Example() {
	fmt.Println("> example of waiting some 2 seconds (almost ) via timer ")
	duration := time.Second * time.Duration(2) // the same as time.Second * 1
	timer := time.NewTimer(duration)

	fmt.Println(" reset timer after 1.5 sec")
	time.Sleep(time.Millisecond * 1500)
	timer.Reset(duration)
	time.Sleep(time.Millisecond * 1500)
	timer.Reset(duration)

	// fmt.Println(" stop timer beforehand ")
	// timer.Stop()

	fmt.Println(" waiting for stop  ")
	<-timer.C
	fmt.Println()

	fmt.Println(">  example of sleeping via channel ")
	time.Sleep(10 * time.Millisecond)

	fmt.Println("waiting via channel")
	select {
	case time := <-time.After(100 * time.Millisecond):
		{
			fmt.Printf("timeout: %#v \n", time)
			goto Next
		}
	}
Next:
}
