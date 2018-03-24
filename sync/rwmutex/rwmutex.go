package rwmutex

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func increaseData(index int, data *map[int]int, rwmutex *sync.RWMutex) {
	time.Sleep(time.Millisecond * 20)
	runtime.Gosched()

	rwmutex.Lock()
	rwmutex.Unlock()

	(*data)[index]++
	fmt.Printf("write: [%v] -> %v \n", index, (*data)[index])

}

func printData(data *map[int]int, rwmutex *sync.RWMutex) {
	rwmutex.RLock()
	rwmutex.RUnlock()

	fmt.Printf("[0]=%v   [1]=%v\n", (*data)[0], (*data)[1])
}

func Example() {
	rwmutex := sync.RWMutex{}
	data := map[int]int{}
	data[0] = 0
	data[1] = 0

	fmt.Println("> write data using RWMutex ( Lock/Unlock ) ")
	for i := 0; i < 10; i++ {
		go increaseData(i%2, &data, &rwmutex)
	}

	fmt.Println("> read data using RWMutex (RLock/RUnlock)")
	for i := 0; i < 10; i++ {
		go printData(&data, &rwmutex)
		time.Sleep(time.Millisecond * 10)
		runtime.Gosched()
	}

	time.Sleep(time.Second * 1)

}
