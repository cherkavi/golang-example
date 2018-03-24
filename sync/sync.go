package main

import (
	"fmt"

	"./mutex"
	"./rwmutex"
)

func main() {
	fmt.Println("---- example of Mutex ----")
	mutex.Example()

	fmt.Println("---- example of RWMutex (read/write mutex) ----")
	rwmutex.Example()
}
