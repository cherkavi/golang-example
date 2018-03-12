package main

import (
	. "fmt"
)

type Point struct {
	Left int
	Right int
}

// new function of struct, read-only access
func (p Point) sum() int {
	return p.Left+p.Right
}

// new function of struct, read-write access
func (p *Point) increase(value int) {
	p.Right += value
	p.Left += value
}

// changes will not be applied
func (p Point) increaseNotAffected(value int) {
	p.Right += value
	p.Left += value
}

func sum(p Point) int{
	return p.Left + p.Right
}

func increase (p *Point, biase int){
	p.Right+=biase
	p.Left+=biase
}

// complex type
type Rectangle struct{
	P1 Point
	P2 Point
}

// embedded type
type Pixel struct{
	Point
	Color string
}

func main() {
	// declaration
	Println(Point{1, 2})
	Println(Point{Right: 2})
	Println(Point{})

	// slice declaration
	Println("slice: ", []Point{
			{2,3},{3,4},{5,6}, {7,8},
		})

	p := Point{1,2}
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

	anonymousSlice:= []struct{a int; b int} {
		{1,2},
		{3,4},
		{4,5},
		{5,6},
	}
	Println("slice of anonymous struct", anonymousSlice)

	// complex value
	r := Rectangle{P1:Point{1,2}, P2:Point{3,4}}
	Println("rectangle: ", r)

	pixel1 := Pixel{}
	pixel1.Left = 3
	pixel1.Right = 5
	pixel1.Color = "red"
	Println("Embedded struct: ", pixel1)

	p1:=Point{1,2}
	p2:=Point{1,2}
	p3:=struct{Left int; Right int}{1,2}
	Printf("point#1:%p %#v  point#2:%p %#v  point#3:%p %#v   point#1==point#2:%v   point#1==point#3:%v", &p1, p1, &p2, p2, &p3, p3, p1==p2, p1==p3, )
}
