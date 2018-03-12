package main

import (
	"fmt"
)

func main(){
	// create map inline
	data:= map[int]string{1:"one", 2:"two"}
	fmt.Printf("%#v    %v \n", data, data)

	// create map via 'make'
	data2:=make(map[string]int, 2) // capacity for two elements only
	data2["one"]=1
	data2["two"]=2
	fmt.Printf("address: %p  len: %v   go: %#v    data: %v \n", data2, len(data2), data2, data2)

	data2["three"]=3 // add extra value - more than capacity
	fmt.Printf("address: %p  len: %v   go: %#v    data: %v \n", data2, len(data2), data2, data2)

}
