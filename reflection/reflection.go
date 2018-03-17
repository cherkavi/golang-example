package main

import (
	"fmt"
	"reflect"

	"./copy"
	"./equals"
	"./struct_analyser"
)

func main() {
	fmt.Println("\n ---------- example of copy data from one struct to another ----")
	copy.ReflectCopy()

	fmt.Println("\n ---------- example print structure code  ----")
	struct_analyser.ReflectInfo(&copy.Movie{"my movie", 100})

	fmt.Println("\n ---------- example of equality  ----")
	equals.Equals()

	fmt.Println("\n ---------- type of element  ----")
	movie := copy.Movie{"Star Wars", 1979}
	fmt.Printf("%#v    type is: %v \n", movie, reflect.TypeOf(movie))
}
