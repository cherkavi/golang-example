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

func (a AStruct) String() string{
	return fmt.Sprintf("override toString method")
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


func whatTheType(i interface{}) {
	switch v := i.(type) {
	//case interface{}: // !!! catch all values !!!
	//	fmt.Println("type is 'interface{}' \n")
	case AStruct:
		fmt.Println("type is AStruct")
	case BStruct:
		fmt.Println("type is BStruct")
	case *BStruct:
		fmt.Println("type is (*reference)BStruct")
	default:
		fmt.Printf("unknown type %T \n", v)
	}
}


func main(){
	fmt.Println("-------- example of implicit implementation for struct and *struct")
	var echo Echo

	echo = AStruct{}
	echo.PrintEcho()
	executeEcho(AStruct{})
	whatTheType(echo)

	echo = &BStruct{} // !!! assignment with pointer, not with value !!!
	echo.PrintEcho()
	executeEcho(&BStruct{})
	whatTheType(echo)

	fmt.Println("\n--------  inline declaration --------")
	var echoInline interface {
		PrintEcho()
	}
	echoInline = echo
	echoInline.PrintEcho()

	fmt.Println("\n-------- empty interface ---------")
	var emptyValue interface {}
	emptyValue = echo
	fmt.Println(emptyValue)

	var emptyValue2 interface {} = "this is just a message"
	fmt.Println(emptyValue2)

	fmt.Println("\n-------- upcasting  ---------")
	var emptyValue3 interface {} = "string from empty interface, upcasting"
	var message string = emptyValue3.(string)
	fmt.Println(message)
	message, convertOk:= emptyValue3.(string)
	if convertOk {
		fmt.Println(message)
	}else{
		fmt.Printf("can't upcast value:%v to string, error: %v \n", message, convertOk)
	}

	fmt.Println("\n-------- upcasting to wrong interface with checking ---------")
	message2, convertOk := emptyValue3.(AStruct)
	if convertOk {
		fmt.Printf("upcast to AStruct %#v\n", message2)
	}else{
		fmt.Printf("can't upcast value:<%v> to AStruct, conversation value: %v \n", emptyValue3, convertOk)
	}

	// fmt.Println("\n-------- upcasting to wrong interface without checking ---------")
	// _ = emptyValue3.(AStruct) // lead to panic

	fmt.Println("\n-------- toString ---------")
	fmt.Println(AStruct{})

	fmt.Println("\n--------  using nil for a pointer --------")
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
