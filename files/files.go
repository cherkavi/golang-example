package main

import (
	"fmt"

	csv "./csv"
	json "./json"
	xml "./xml"
	yaml "./yaml"
)

func main() {
	fmt.Println("---- working with JSON ----")
	json.Example()
	fmt.Println()

	fmt.Println("---- working with XML ----")
	xml.Example()
	fmt.Println()

	fmt.Println("---- working with Yaml ----")
	yaml.Example()
	fmt.Println()

	fmt.Println("---- working with CSV ----")
	csv.Example()
	fmt.Println()
}
