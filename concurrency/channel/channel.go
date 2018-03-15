package channel

import (
	"time"
	"fmt"
)


type Flag struct{
	name string
	value bool
}

func (f Flag) String() string{
	return fmt.Sprintf("-[%v=%v]-", f.name, f.value)
}

func pingPong(marker string, channel chan Flag){
	for i:=0; i<5; i++{
		fmt.Printf("waiting...%v", marker)
		flag:= <- channel
		flag.value=!flag.value
		fmt.Printf(" %v\n", flag)
		channel<-flag
		time.Sleep(time.Millisecond*250)
	}
}

func PingPong(){
	sharedChannel:=make(chan Flag)
	go pingPong("1", sharedChannel)
	go pingPong("2", sharedChannel)
	sharedChannel <- Flag{name:"ball", value:true}
	time.Sleep(time.Second*5)
	close(sharedChannel)
}