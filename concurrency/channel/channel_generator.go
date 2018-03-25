package channel

import (
	"fmt"
	"time"
)

func process() chan<- bool {
	control := make(chan bool)
	go func() {
		for {
			select {
			case controlValue := <-control:
				{
					if controlValue {
						fmt.Println("signal received: start ")
					} else {
						fmt.Println("signal received: stop ")
						break
					}
				}
			}
		}
	}()
	return control
}

func integerGenerator(values ...int) <-chan int {
	returnValue := make(chan int)
	go func() {
		var value int
		for _, value = range values {
			returnValue <- value
			time.Sleep(time.Millisecond * 100)
		}
		close(returnValue)
	}()
	return returnValue
}

func ExampleOfChannelGenerator() {
	fmt.Println("> execute process and receive channel as a control knob  ")
	controlKnob := process()
	controlKnob <- true
	controlKnob <- false

	fmt.Println("> print range of digits ")
	for i := range integerGenerator(10, 11, 15, 8, 7, 5) {
		fmt.Printf("%v  ", i)
	}
	fmt.Println()
}
