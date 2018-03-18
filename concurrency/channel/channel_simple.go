package channel

import (
	"fmt"
	"time"
)

func waitAndPrint(signal <-chan interface{}, message string) {
	<-signal // just waiting for a signal and than - print message
	fmt.Println(message)
}

func waitingForEmptyChannel(notifier chan interface{}) {
	for {
		if len(notifier) == 0 {
			goto End // can be changed to return
		} else {
			time.Sleep(time.Millisecond * 10)
			fmt.Print("*")
		}
	}
End:
}

func SpreadSignals() {
	notifier := make(chan interface{}, 10)
	messages := []string{" this is ", " message ", " from ", " goroutines "}

	// execute respective amount of goroutines
	for _, eachMessage := range messages {
		go waitAndPrint(notifier, eachMessage)
	}
	// notify each of them about printing
	for index := range messages {
		notifier <- index // just send message into channel
	}

	waitingForEmptyChannel(notifier)
	fmt.Println()
}
