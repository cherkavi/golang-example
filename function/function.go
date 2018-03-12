package main

import (
	"fmt"
	// example of inline import - don't need to specify name of the package
	. "./inner"

	// example of import with alias
	f "./inner2"

	// example of import with side-effect ( function "init" will be executed )
	_ "./inner3"
)

func privateFunctionWithoutReturnValue() {
}

func functionSwap(x, y string) (string, string) {
	return y, x
}

func privateFunctionSwap2(x, y string) (b, a string) {
	a = x
	b = y
	return
}

func closureExample(magicNumber int) func(koef int) int {
	magicNumber += 2
	var offset = magicNumber * 2
	return func(a int) int {
		return offset + a*2
	}
}

func executeFunction(f func(a int) int) int {
	return f(5)
}

func main() {
	privateFunctionWithoutReturnValue()

	x, y := functionSwap("1", "2")
	x, y = privateFunctionSwap2("1", "2")
	fmt.Printf("local swap function: %v   %v \n", x, y)

	sum := ExternalAdd(1, 2, 3, 4)
	fmt.Println("execute external function ", sum)

	array := []int{1, 2, 3, 4}
	mul := f.ExternalMultiplication(array...)
	fmt.Println("execute external function via alias", mul)

	fmt.Println("closure example: ", closureExample(5)(2))

	var closureFunction = closureExample(512)
	closureFunction(5)

	var closureFunction2 func(b int) int
	closureFunction2 = closureExample(7)
	closureFunction2(23)

	a := executeFunction(closureFunction2)
	fmt.Println(" result from function where parameter was another function ", a)

}