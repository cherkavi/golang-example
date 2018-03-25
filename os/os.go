package main

import (
	"fmt"

	"./signal"
	"./workingwithfile"
)

func main() {
	fmt.Println("---- working with files ---- ")
	workingwithfile.FileOperations()
	fmt.Println()

	fmt.Println("---- catch signals from OS  ---- ")
	signal.ExampleOfSignalHook()
}
