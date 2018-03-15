package main

import (
	"fmt"
	"./simple"
	"./channel"
)


func main(){
	fmt.Println("--- parallel threads execution ---")
	simple.ExampleOfSimpleExecution()

	fmt.Println("--- parallel threads execution ---")
	channel.PingPong()

	fmt.Println("--- using range to read from channel ---")
	channel.RangeAllFromChannel()

	fmt.Println("--- waiting for signal from multiply sources ---")
	channel.WaitingAndReadingFromMultiplySources()

	fmt.Println("--- using Mutex (Lock) ---")
	channel.MutexUsing()
}
