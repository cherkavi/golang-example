package once

import (
	"fmt"
	"sync"
	"time"
)

var once = sync.Once{}

var onceFunc = func() {
	fmt.Println("I'm executing only once")
}

func Example() {
	fmt.Println("> function execution will appear once only ")
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceFunc)
		}()
	}
	time.Sleep(time.Second * 2)
}
