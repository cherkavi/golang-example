package main

import (
	"fmt"
	"strconv"
)


type NewInteger int64

func (value NewInteger) Len() int{ // copy structure each time when it executed
	return len(strconv.Itoa(int(value)))
}

func (value *NewInteger) increment() int { // just send via pointer
	var newValue = int(*value)+1
	*value = NewInteger(newValue)
	return int(*value)
}

func reset(value *NewInteger){
	*value = NewInteger(0)
}

func main(){
	var value int = 15
	var newValue NewInteger = NewInteger(value)
	fmt.Printf("value: %v    len: %v \n", value, newValue)
	fmt.Printf("value: %v    increase: %v \n", value, newValue.increment())

	reset(&newValue) // using as a Pointer to send value
	fmt.Printf("value: %v    increase after reset: %v \n", value, newValue.increment())

}
