package channel

import (
	"fmt"
	"math/rand"
	"time"
)

// waiting first data from one of the channel
// channels declared as read-only
func waitingForFirstSignal(a, b, c <-chan string) {
	var data string
	select { // with "default" section - non-blocking waiting
	case data = <-a:
		{
			fmt.Printf("from channel 'a': %v\n", data)
		}
	case data = <-b:
		{
			fmt.Printf("from channel 'b': %v\n", data)
		}
	case data = <-c: // here also can be writer to another channel, not reader only  ( with allowing to write )
		{
			fmt.Printf("from channel 'c': %v\n", data)
		}
	case <-time.After(20 * time.Second): // timeout after 20 seconds
		{
			fmt.Println("!!!Timeout!!!")
		}
		// uncomment next block and "select" will not wait
		//default: {
		//	fmt.Println("no wait, no data from channels")
		//}
	}
}

func pushDataRandomly(channels ...chan<- string) {
	channels[rand.Intn(len(channels))] <- time.Now().Format(time.Kitchen)
}

func WaitingAndReadingFromMultiplySources() {
	a := make(chan string)
	b := make(chan string)
	c := make(chan string)

	go waitingForFirstSignal(a, b, c)
	go pushDataRandomly(a, b, c)

	time.Sleep(time.Second * 1)
	fmt.Println()
}
