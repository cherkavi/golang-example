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
	time.Sleep(time.Millisecond * 10)
}

func SpreadSignals() {
	messages := []string{" this ", " is ", " message ", " from ", " goroutines ", " and ", " for ", " console "}
	notifier := make(chan interface{}, len(messages)/3) // buffer is less than queue

	// execute respective amount of goroutines
	for _, eachMessage := range messages {
		go waitAndPrint(notifier, eachMessage)
	}
	// notify each of them about printing
	for index := range messages {
		notifier <- index // just send message into channel, wait when buffer is full
		fmt.Print(">")
	}

	waitingForEmptyChannel(notifier)
	fmt.Println()
}
