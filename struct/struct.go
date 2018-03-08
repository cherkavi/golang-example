package main

import (
	. "fmt"
)

type Pair struct {
	Left int
	Right int
}

// new function of struct, read-only access
func (p Pair) sum() int {
	return p.Left+p.Right
}

// new function of struct, read-write access
func (p *Pair) increase(value int) {
	p.Right += value
	p.Left += value
}

// changes will not be applied
func (p Pair) increaseNotAffected(value int) {
	p.Right += value
	p.Left += value
}

func sum(p Pair) int{
	return p.Left + p.Right
}

func increase (p *Pair, biase int){
	p.Right+=biase
	p.Left+=biase
}

func main() {
	// declaration
	Println(Pair{1, 2})
	Println(Pair{Right: 2})
	Println(Pair{})

	// slice declaration
	Println("slice: ", []Pair{
			{2,3},{3,4},{5,6}, {7,8},
		})

	p := Pair{1,2}
	Println("value:", p)
	Println("sum:", p.sum(), "  ", sum(p))

	p.increaseNotAffected(5)
	Println("increase (not affected):", p)

	p.increase(5)
	Println("increase:", p)

	increase(&p, 7)
	Println("increase via function:", p)

	noNameStruct := struct{i int; s string} {i:10, s: "error_message"}
	Println("anonymous struct: ",  noNameStruct )
}
