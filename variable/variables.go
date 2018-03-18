package main

import (
	"fmt"

	"./constants"
)

func changeOuterVariable(external *int) {
	*external = *external + 22
}

func applyFunction(customFunction func(x int) int, argument int) int {
	return customFunction(argument)
}

const privateVisibleInt uint64 = 9999999999999999

// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr

// byte
// rune - unicode representation

// float32, float64

// complex64, complex128

func main() {
	// declare variable with default value
	var a int
	fmt.Printf("default value for integer is %v \n", a)

	// declare variable and init it
	var b int = 2
	fmt.Printf("declare and init value  %v\n", b)
	// multiply assign
	var x, y int = 10, 15
	fmt.Println("multiply assignment: ", x, y)

	// declare variable and not to mention about type ( automatic type assign )
	c := 2
	// multiply return value from function
	bytesNumber, error := fmt.Printf("just init value: %v and assign type automatically: %T\n", c, c)

	// multiply assignment
	w, z := 10, 15
	fmt.Println("multiply assignment with automatic type: ", w, z)

	if error != nil {
		fmt.Printf(" %v amount of bytes were written, with error %v", bytesNumber, error)
	}

	fmt.Println("value from another package: ", constants.A)

	fmt.Println("const from another package: ", constants.Message)

	d := 10
	fmt.Println("variable before changing : ", d)
	changeOuterVariable(&d)
	fmt.Println("variable after changing : ", d)

	var arrayExample = [5]uint64{10, 11, 12, 13, 14}
	fmt.Printf("Array declaration: %#v\n", arrayExample)

	var sliceExample = []uint64{10, 11, 12, 13}
	fmt.Printf("Slice declaration: %#v\n", sliceExample)

	var customFunction = func(x int) int {
		return x + 1
	}

	for _, eachValue := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("%v ", applyFunction(customFunction, eachValue))
	}

}
