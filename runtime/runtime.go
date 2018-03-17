package main

import (
	"fmt"
	"runtime"
)

func printOS() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	default:
		// freebsd, openbsd,plan9
		fmt.Printf("%s \n", os)
	}
}

func runGarbageCollector() {
	runtime.GC()
	fmt.Println("Garbage Collector was started")
}

func amountOfInvolvedProcessors() {
	fmt.Printf("amount of processors that can be involved: %v \n", runtime.GOMAXPROCS(0))
}

func amountOfProcessors() {
	fmt.Printf("amount of processors: %v \n", runtime.NumCPU())
}

func printGoRoot() {
	fmt.Printf("GOROOT: %v \n", runtime.GOROOT())
}

func main() {
	printOS()
	runGarbageCollector()
	amountOfProcessors()
	amountOfInvolvedProcessors()
	printGoRoot()

	// runtime.Goexit() // can be called not for main
}
