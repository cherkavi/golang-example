package method_expression

import (
	"fmt"
)

type IPlay interface {
	Play()
}

type A struct {
	name string
}

func (a A) String() string {
	return a.name
}

func (a A) Play() {
	fmt.Printf("%#v\n", a)
}

type B struct {
}

func (b B) Play() {
	fmt.Printf("%#v\n", b)
}

func MethodExpression() {
	fmt.Println("---- example of using interface ----  ")
	var musicTools []IPlay = []IPlay{A{}, B{}}
	for _, eachTool := range musicTools {
		eachTool.Play()
	}

	fmt.Println("---- example of using pointer to method ---- ")
	pointerToMethod := A.Play
	// fmt.Printf("type of pointer: %#v\n", pointerToMethod)
	for _, each := range []A{A{"name1"}, A{"name2"}, A{"name3"}, A{"name4"}} {
		// execute method against each element
		pointerToMethod(each)
	}

}
