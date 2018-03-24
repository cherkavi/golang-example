package main

import (
	"fmt"
)

func execute() {
	fmt.Println("begin of function ")
	defer fmt.Println(" defer 1")
	defer fmt.Println(" defer 2")
	defer fmt.Println(" defer 3")

	fmt.Println("end of function ")
}

func main() {
	fmt.Println("---- example of ordering of defer ")
	execute()
}
