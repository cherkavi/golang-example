package main

import (
	"fmt"
	"time"
	"./simple"
	"./channel"
)


func main(){
	fmt.Println("--- parallel threads execution ---")
	simple.ExampleOfSimpleExecution()
	time.Sleep(time.Second*3)

	fmt.Println("--- parallel threads execution ---")
	channel.PingPong()
}
