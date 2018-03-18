package simple

import (
	"fmt"
	"time"
)

func delayPrint(message string) {
	time.Sleep(time.Second * 1)
	fmt.Println("message from goroutine: " + message)
}

func delayExecution(externalFunction func()) {
	time.Sleep(time.Second * 1)
	externalFunction()
}

func ExecuteWithParameter() {
	fmt.Printf("\n ---- attempt to print message after change it ---- \n")
	message := "first message"
	fmt.Printf("message before start goroutine: %v \n", message)
	go delayPrint(message)
	message = "changed message"
	fmt.Printf("message after start goroutine: %v \n", message)
	time.Sleep(time.Second * 2)

	fmt.Printf("\n ---- attempt to execute function into closure after change environment ---- \n")
	message = "first message"
	fmt.Printf("message before start goroutine: %v \n", message)
	go delayExecution(func() {
		fmt.Println("message from goroutine: " + message)
	})
	message = "changed message"
	fmt.Printf("message after start goroutine: %v \n", message)
	time.Sleep(time.Second * 2)
	fmt.Println()
}
