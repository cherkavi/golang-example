package main

import (
	"fmt"

	"./subpackage"
)

func init() {
	fmt.Println("init method of main package ")
}

func main() {
	fmt.Println("main method of main package was started ")
	subpackage.JustEmptyFunction()
}
