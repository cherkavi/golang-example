package main

import (
	"fmt"

	"./subpackage"
)

func init() {
	fmt.Println("init method of main package ")
}

func init() {
	fmt.Println("init method ( second ) of main package ")
}

func init() {
	fmt.Println("init method ( third ) of main package ")
}

func main() {
	fmt.Println("main method of main package was started ")
	subpackage.JustEmptyFunction()
}
