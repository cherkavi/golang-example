package main

import (
	"fmt"
	"runtime"
	"time"
)

func finalizeExample() {
	var array = []int{1, 2, 3, 4, 5}
	fmt.Printf(" Array: %v \n", array)

	runtime.SetFinalizer(&array, func(array *[]int) {
		fmt.Printf("finalize for array: %v \n", array)
	})
}

func main() {
	finalizeExample()
	fmt.Println("execute Garbage collector:")
	time.Sleep(time.Second * 1)
	runtime.GC()
	time.Sleep(time.Second * 1)
}
