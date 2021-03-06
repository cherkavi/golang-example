package main

import (
	"fmt"

	"./channel"
	"./simple"
)

func main() {
	//A send to a nil channel blocks forever
	//A receive from a nil channel blocks forever
	//A send to a closed channel panics
	//A receive from a closed channel returns default init value immediately

	fmt.Println("--- parallel threads execution ---")
	simple.ExampleOfSimpleExecution()

	fmt.Println("--- parallel threads execution ---")
	channel.PingPong()

	fmt.Println("--- example of timer based on select ---")
	channel.ExampleSelectTimer()

	fmt.Println("--- pattern ChannelGenerator ---")
	channel.ExampleOfChannelGenerator()

	fmt.Println("--- pattern ChannelPipeline ---")
	channel.ExampleChannelPipeline()

	fmt.Println("--- merge channels ---")
	channel.ExampleOfMerge()

	fmt.Println("--- using range to read from channel ---")
	channel.RangeAllFromChannel()

	fmt.Println("--- waiting for signal from multiply sources ---")
	channel.WaitingAndReadingFromMultiplySources()

	fmt.Println("--- using Mutex (Lock) ---")
	channel.MutexUsing()

	fmt.Println("--- goroutines with parameters ---")
	simple.ExecuteWithParameter()

	fmt.Println("--- using channel as notificator ---")
	channel.SpreadSignals()
}
