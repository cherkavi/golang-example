package main

import (
	"fmt"
	"reflect"
)

type Point struct{
	X, Y int
}

func main(){
	// create map inline
	data:= map[int]string{1:"one", 2:"two"}
	fmt.Printf("%#v    %v \n", data, data)

	// create map via 'make'
	data2:=make(map[int]string, 2) // capacity for two elements only
	data2[1]="one"
	data2[2]="two"
	fmt.Printf("address: %p  len: %v   go: %#v    data: %v \n", data2, len(data2), data2, data2)

	data2[3]="three" // add extra value - more than capacity
	fmt.Printf("address: %p  len: %v   go: %#v    data: %v \n", data2, len(data2), data2, data2)
	delete(data2, 3) // delete data

	if _, exists := data2[3]; !exists{
		fmt.Printf("data was removed: '3' \n")
	}

	fmt.Printf("%#v   is equals to %#v  - %v\n", data, data2, reflect.DeepEqual(data, data2))

	rectangle := map[string]Point {
		"upper-left":{10,20},
		"upper-right":{20,20},
		"bottom-right":{20,10},
		"bottom-left":{10,10},
	}
	for key,value := range rectangle{
		fmt.Printf("key=%v   value=%v\n", key, value)
	}

}
