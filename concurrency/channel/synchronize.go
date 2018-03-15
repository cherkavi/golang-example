package channel

import (
	"sync"
	"runtime"
	"bytes"
	"strconv"
	"fmt"
	"time"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

type Bucket struct{
	value int
}

func increaseBucket(mutex *sync.Mutex, bucket *Bucket){
	mutex.Lock()
	defer mutex.Unlock()

	bucket.value+=1
	fmt.Printf("Thread:%v   Bucket.Value: %v \n", getGID(), bucket.value)
}


func increaseBucketTimes(mutex *sync.Mutex, bucket *Bucket, times int){
	for index:=0;index<times;index++{
		increaseBucket(mutex, bucket)
		time.Sleep(time.Millisecond*150) // move it inside method above
	}
}

func MutexUsing(){
	bucket:=Bucket{0}
	mutex:=sync.Mutex{}

	// mutex and bucket must be sent into procedure using pointers
	go increaseBucketTimes(&mutex, &bucket, 15)
	go increaseBucketTimes(&mutex, &bucket, 15)
	go increaseBucketTimes(&mutex, &bucket, 15)

	time.Sleep(time.Second*4)
}