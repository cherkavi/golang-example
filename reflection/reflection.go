package main

import (
	"fmt"
	"reflect"

	"./copy"
	"./equals"
	"./struct_analyser"
)

func instanceof(v interface{}) { // input parameter - empty interface
	typeOf := reflect.TypeOf(v)
	if typeOf.Kind() == reflect.Int {
		fmt.Println("integer type detected")
	}
	fmt.Printf("typeOf: [%#v] is [%v] \n", v, typeOf)

	// first law of reflection: reflect.ValueOf(v)
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() == reflect.Int {
		// second law of reflection: valueOf.Interface()
		fmt.Println("integer value detected: ", valueOf.Interface().(int)) // convert to real value and upcast
	}
	// third law of reflection: valueOf.CanSet()
	fmt.Printf("valueOf: [%#v] is [%v] by type [%v] is changable: [%v] \n", v, valueOf, valueOf.Type(), valueOf.CanSet())
	// fmt.Println("Elements: ", valueOf.Elem())
	// see patterns/creational/builder.go

	fmt.Println()
}

func isValidPointer(v interface{}) bool {
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() != reflect.Ptr || valueOf.IsNil() {
		return false
	} else {
		// it is not nil and pointer
		return true
	}
}

// see https://blog.golang.org/laws-of-reflection
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

	fmt.Println("\n ---------- using any type of value to print -----")
	instanceof(10)
	instanceof("this is string")
	s := "this is string pointer"
	instanceof(&s)
	instanceof(func() {})
	instanceof(struct{}{}) // example of empty struct
}
