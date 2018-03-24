package mutex

import (
	"fmt"
	"sync"
	"time"
)

func increaseWithoutLock(data *int32) {
	*data++
	fmt.Printf("%v\n", *data)
	time.Sleep(time.Millisecond * 200)
}

func increaseWithLock(data *int32, mutex sync.Locker) {
	mutex.Lock()
	defer mutex.Unlock()

	increaseWithoutLock(data)
}

func Example() {
	var data int32 = 0
	fmt.Println("> execute without lock ")
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				increaseWithoutLock(&data)
			}
		}()
	}
	time.Sleep(time.Second * 2)

	data = 0
	mutex := sync.Mutex{} // implementation of interface sync.Locker
	fmt.Println("> execute with lock ")
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				increaseWithLock(&data, &mutex)
			}
		}()
	}
	time.Sleep(time.Second * 2)
}
