package channel

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// Start an output goroutine for each input channel in cs.  output
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Fan In - Fan Out
func ExampleOfMerge() {
	fmt.Println("---- merge data from different channels ")

	fmt.Println(" list1: 5,6,7,8,9")
	list1 := integerGenerator(5, 6, 7, 8, 9)
	fmt.Println(" list1: 1,2,3,4,")
	list2 := integerGenerator(1, 2, 3, 4)

	for value := range merge(list1, list2) {
		fmt.Printf("%v ", value)
	}
	fmt.Println()
}
