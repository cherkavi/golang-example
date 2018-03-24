package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func Example() {
	wg := sync.WaitGroup{}

	fmt.Println("> start threads ")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Printf("next routine was started: %v\n", index)
			time.Sleep(time.Millisecond * 50)
			wg.Done()
		}(i)
	}
	fmt.Println("> all thread were initiated, waiting for finish")
	wg.Wait()
	fmt.Println("> done")
}
