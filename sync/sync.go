package main

import (
	"fmt"

	"./mutex"
	"./once"
	"./rwmutex"
)

func main() {
	fmt.Println("---- example of Mutex ----")
	mutex.Example()

	fmt.Println("---- example of RWMutex (read/write mutex) ----")
	rwmutex.Example()

	fmt.Println("---- example of Once (execute func only once) ----")
	once.Example()
}
