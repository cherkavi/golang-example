package main

import (
	"fmt"

	json "./json"
	xml "./xml"
)

func main() {
	fmt.Println("---- working with JSON ----")
	json.Example()
	fmt.Println()

	fmt.Println("---- working with XML ----")
	xml.Example()
}
