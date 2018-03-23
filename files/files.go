package main

import (
	"fmt"

	csv "./csv"
	json "./json"
	xml "./xml"
)

func main() {
	fmt.Println("---- working with JSON ----")
	json.Example()
	fmt.Println()

	fmt.Println("---- working with XML ----")
	xml.Example()
	fmt.Println()

	fmt.Println("---- working with CSV ----")
	csv.Example()
	fmt.Println()
}
