package main

import (
	"fmt"
	"time"
)

type Echo interface{
	// just print output info
	PrintEcho()
}


type AStruct struct{
}
func (a AStruct) PrintEcho(){ // no 'implements' declaration need
	fmt.Printf("%#v \n", a)
}


type BStruct struct{
}
func (b *BStruct) PrintEcho(){ // !!! pointer implemented interface !!!
	if b==nil{
		fmt.Println("!!!nil!!!")
		return
	}
	fmt.Printf("%#v \n", b)
}


func executeEcho(e Echo){
	e.PrintEcho()
}

func main(){
	fmt.Println("example of implicit implementation for struct and *struct")
	var echo Echo

	echo = AStruct{}
	echo.PrintEcho()
	executeEcho(AStruct{})

	echo = &BStruct{} // !!! assignment with pointer, not with value !!!
	echo.PrintEcho()
	executeEcho(&BStruct{})

	fmt.Println("--------  inline declaration --------")
	var echoInline interface {
		PrintEcho()
	}
	echoInline = echo
	echoInline.PrintEcho()

	fmt.Println("-------- empty interface ---------")
	var emptyValue interface {}
	emptyValue = echo
	fmt.Println(emptyValue)

	var emptyValue2 interface {} = "this is just a message"
	fmt.Println(emptyValue2)

	fmt.Println("-------- upcasting  ---------")
	var emptyValue3 interface {} = "string from empty interface, upcasting"
	var message string = emptyValue3.(string)
	fmt.Println(message)


	fmt.Println("--------  using nil for a pointer --------")
	var b2 BStruct
	echo = &b2
	echo.PrintEcho()

	var b3 *BStruct
	echo = b3
	echo.PrintEcho()

	time.Sleep(time.Second*2)
	var e2 Echo
	e2.PrintEcho() // attempt to call value without implementation

}
