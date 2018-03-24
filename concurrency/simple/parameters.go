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

type Data struct {
	value string
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

	fmt.Printf("\n ---- operate with closure data into go-routine ( using value from parent Thread) ----")
	fmt.Printf("\n ---- object will have final value from Parent Thread ( 4 ) and address of object will be the same  ---- \n")
	dataFromMainThread := Data{}
	for i := 0; i < 5; i++ {
		dataFromMainThread.value = fmt.Sprintf("closure data: %v", i)
		go func() {
			time.Sleep(time.Millisecond * 100)
			for i := 0; i < 5; i++ {
				fmt.Printf("%v  %#v  %p \n", dataFromMainThread, dataFromMainThread, &dataFromMainThread)
				// time.Sleep(time.Millisecond * 50)
			}
		}()
	}
	time.Sleep(time.Second * 1)

	fmt.Printf("\n ---- operate with copy of data into go-routine ( from parent Thread) ---- ")
	fmt.Printf("\n ---- different values, different addresses of object ---- \n")
	for i := 0; i < 5; i++ {
		dataFromMainThread.value = fmt.Sprintf("copy of data: %v", i)
		go func(dataFromMainThread Data) {
			time.Sleep(time.Millisecond * 100)
			for i := 0; i < 5; i++ {
				fmt.Printf("%v  %#v  %p \n", dataFromMainThread, dataFromMainThread, &dataFromMainThread)
				// time.Sleep(time.Millisecond * 50)
			}
		}(dataFromMainThread)
	}
	fmt.Println("    parent generator finished")
	time.Sleep(time.Second * 1)
	fmt.Println()

}
