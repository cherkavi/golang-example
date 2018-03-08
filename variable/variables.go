package main

import (
	"./constants"
	"fmt"
)


func changeOuterVariable(external *int){
	*external= *external + 22
}

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
	changeOuterVariable(&d)
	fmt.Println("variable after changing : ", d)
}
