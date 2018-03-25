package main

import (
	"fmt"

	"./signal"
	"./std"
	"./workingwithfile"
)

func main() {
	fmt.Println("---- working with files ---- ")
	workingwithfile.FileOperations()
	fmt.Println()

	fmt.Println("---- read data from console ")
	std.ExampleOfReadingText()

	fmt.Println("---- catch signals from OS  ---- ")
	signal.ExampleOfSignalHook()

}
