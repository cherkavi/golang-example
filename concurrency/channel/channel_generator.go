package channel

import (
	"fmt"
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

func ExampleOfChannelGenerator() {
	fmt.Println("> execute process and receive channel as a control knob  ")
	controlKnob := process()
	controlKnob <- true
	controlKnob <- false
}
