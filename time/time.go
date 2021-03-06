package main

import (
	"fmt"

	"./format"
	"./timer"
)

func main() {
	fmt.Println("---- example of print data/time with format ")
	format.Example()

	fmt.Println("---- example of using timer ")
	timer.Example()

	fmt.Println("---- example of using ticker ")
	timer.ExampleTicker()
}
