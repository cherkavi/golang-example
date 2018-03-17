package method_expression

import (
	"fmt"
	"reflect"
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
	objectA := A{}
	objectB := B{}
	fmt.Println("---- example of using interface ----  ")
	var musicTools []IPlay = []IPlay{objectA, objectB}
	for _, eachTool := range musicTools {
		eachTool.Play()
	}

	pointerA := &A{}
	pointerPlay := pointerA.Play
	fmt.Printf("type of pointer to Play method: %#v\n", reflect.TypeOf(pointerPlay))
	fmt.Printf("value of pointer to Play method %v \n", pointerPlay)
	fmt.Printf("value of pointer to object: %v \n", pointerA)

	pointerToMethod := A.Play
	// fmt.Printf("type of pointer: %#v\n", pointerToMethod)
	for _, each := range []A{A{"name1"}, A{"name2"}, A{"name3"}, A{"name4"}} {
		// execute method against each element
		pointerToMethod(each)
	}

}
