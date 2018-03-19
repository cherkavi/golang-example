package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var stringParam string
var intParam int

type arrayOfInt []int

var listParam arrayOfInt

func (array *arrayOfInt) Set(inputValue string) error {
	for _, value := range strings.Fields(inputValue) {
		if intValue, exception := strconv.Atoi(value); exception == nil {
			*array = append(*array, intValue)
		}
	}
	return nil
}

func (array arrayOfInt) String() string {
	if array == nil {
		return ""
	}
	var returnValue strings.Builder
	for index := 0; index < len(array); index++ {
		returnValue.WriteString(strconv.Itoa(array[index]))
	}
	return fmt.Sprintf("%v", returnValue.String())
}

func init() {
	flag.StringVar(&stringParam, "stringParam", "default string", "just a string value")
	flag.IntVar(&intParam, "intParam", 5, "just a int value")
	flag.Var(&listParam, "values", "just a string value")
	// flag.Parse() --- not here !!!!
}

// go run flag/flag.go -stringParam="hello" -intParam=7 -values "1 2 3 4"
func main() {
	flag.Parse() // should be executed into "main" method
	fmt.Println(stringParam)
	fmt.Println(intParam)
	fmt.Printf("%#v\n", listParam)
}
