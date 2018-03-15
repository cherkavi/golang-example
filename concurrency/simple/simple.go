package simple

import (
	"fmt"
	"time"
)


func counter(marker string){
	for i:=0; i<10; i++{
		fmt.Printf("%v", marker)
		time.Sleep(time.Millisecond*250)
	}
}

func ExampleOfSimpleExecution(){
	go counter("1.")
	go counter("2.")
	go counter("3.")

	time.Sleep(time.Second*3)
	fmt.Println()
}

