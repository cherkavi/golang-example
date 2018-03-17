package channel

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func getGID() uint64 { // runtime.NumGoroutine() will not return number of current goroutine
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

type Bucket struct {
	value int
}

func increaseBucket(mutex *sync.Mutex, bucket *Bucket) {
	mutex.Lock()
	defer func() {
		mutex.Unlock()
		runtime.Gosched() // equals Thread.currentThread().yield()
	}()

	bucket.value++
	fmt.Printf("Thread:%v   Bucket.Value: %v \n", getGID(), bucket.value)

	time.Sleep(time.Millisecond * 150) // move it inside method above
}

func increaseBucketTimes(mutex *sync.Mutex, bucket *Bucket, times int) {
	for index := 0; index < times; index++ {
		increaseBucket(mutex, bucket)
	}
}

func MutexUsing() {
	bucket := Bucket{0}
	mutex := sync.Mutex{}

	// mutex and bucket must be sent into procedure using pointers
	go increaseBucketTimes(&mutex, &bucket, 15)
	go increaseBucketTimes(&mutex, &bucket, 15)
	go increaseBucketTimes(&mutex, &bucket, 15)

	time.Sleep(time.Second * 4)
}
