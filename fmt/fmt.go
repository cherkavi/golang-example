package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("---- example of printing values ----")
	array := []int{5, 6, 7, 8, 9}
	fmt.Printf(" value: %v   go-representation: %#v    char-representation:  %q  \n", array, array, array)

	fmt.Println("---- example of parsing values ----")
	stringForParsing := " 2018 Austria 5.45 "
	var year int
	var countryName string
	var value float32
	fmt.Fscanf(strings.NewReader(stringForParsing), " %v %v %v", &year, &countryName, &value)
	fmt.Printf("Year: %v   Country: %v   Value: %v\n", year, countryName, value)
}
