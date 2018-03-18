package channel

import (
	"fmt"
	"time"
)

type Flag struct {
	name  string
	value bool
}

func (f Flag) String() string {
	return fmt.Sprintf("-[%v=%v]-", f.name, f.value)
}

func pingPong(marker string, channel chan Flag) {
	// catch exception if it will be
	defer func() {
		if errorMessage := recover(); errorMessage != nil {
			fmt.Printf("catch error: [%v] \n", errorMessage)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Printf("waiting...%v", marker)
		flag, readOk := <-channel //
		if !(readOk) {
			fmt.Printf("%v read impossible - channel was closed \n", marker)
			return
		}
		flag.value = !flag.value

		// example of non-blocking sending ( the same for reading )
	label:
		select {
		case channel <- flag:
			fmt.Printf(" %v\n", flag)
		default:
			// can't send - waiting
			time.Sleep(time.Millisecond * 50)
			fmt.Print("*")
			goto label
		}
	}
}

func PingPong() {
	// readonly  channel: make(<-chan Flag)
	// writeonly channel: make(chan<- Flag)
	sharedChannel := make(chan Flag)
	go pingPong("1", sharedChannel)
	go pingPong("2", sharedChannel)
	sharedChannel <- Flag{name: "ball", value: true}
	time.Sleep(time.Second * 1)
	close(sharedChannel)
	time.Sleep(time.Second * 2)
	fmt.Println()
}
