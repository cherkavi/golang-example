package channel

import (
	"fmt"
	"time"
)

// using usual channel as writeonly channel
func writeIn(destination chan<- string) {
	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("%v", i)
		fmt.Printf(" w->%v- ", data)
		destination <- data
	}
	close(destination)
}

// using usual channel as readonly channel
func readFrom(source <-chan string) {
	for msg := range source {
		fmt.Printf(" %v->r ", msg)
	}
}

func RangeAllFromChannel() {
	channel := make(chan string, 3)
	go writeIn(channel)
	go readFrom(channel)
	time.Sleep(time.Second * 1)
	fmt.Println()
	fmt.Println()
}
