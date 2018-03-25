package channel

import (
	"fmt"
	"time"
)

func executeTimer(duration time.Duration, stopChannel <-chan struct{}, resetChannel <-chan time.Duration) {
	timerDuration := duration
MainLoop:
	for {
		select {
		case <-time.After(duration):
			fmt.Println("timer ")
		case <-stopChannel:
			fmt.Println("timer stop")
			break MainLoop // just a break will exit from "select" block
		case timerDuration = <-resetChannel:
			fmt.Printf("timer was changed to : %v \n", timerDuration)
			continue
		}
	}
}

func ExampleSelectTimer() {
	stopChannel := make(chan struct{})
	resetChannel := make(chan time.Duration)

	fmt.Println("> execute controlled via channels timer ")
	go executeTimer(time.Second*1, stopChannel, resetChannel)

	time.Sleep(time.Second * 3)
	fmt.Println(" reset timer: ")
	resetChannel <- time.Millisecond * 500

	time.Sleep(time.Second * 2)
	fmt.Println(" stop timer: ")
	stopChannel <- struct{}{}

}
