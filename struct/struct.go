package main

import (
	. "fmt"
)

type Point struct {
	Left  int
	Right int
}

func newPoint() *Point {
	return new(Point) // return &Point{}
}

// new function of struct, read-only access
func (p Point) sum() int {
	return p.Left + p.Right
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

func sum(p Point) int {
	return p.Left + p.Right
}

func increase(p *Point, biase int) {
	p.Right += biase
	p.Left += biase
}

// complex type
type Rectangle struct {
	P1 Point    // regular struct
	P2 struct { // nested struct
		Left  int
		Right int
	}
}

// embedded type
type Pixel struct {
	Point // embedded struct
	Color string
}

func main() {
	// declaration
	Println(Point{1, 2})
	Println(Point{Right: 2})
	Println(Point{})

	// slice declaration
	Println("slice: ", []Point{
		{2, 3}, {3, 4}, {5, 6}, {7, 8},
	})

	p := Point{1, 2}
	Println("value:", p)
	Println("sum:", p.sum(), "  ", sum(p))

	p.increaseNotAffected(5)
	Println("increase (not affected):", p)

	p.increase(5)
	Println("increase:", p)

	increase(&p, 7)
	Println("increase via function:", p)

	noNameStruct := struct {
		i int
		s string
	}{i: 10, s: "error_message"}
	Println("anonymous struct: ", noNameStruct)

	anonymousSlice := []struct {
		a int
		b int
	}{
		{1, 2},
		{3, 4},
		{4, 5},
		{5, 6},
	}
	Println("slice of anonymous struct", anonymousSlice)

	// complex value
	r := Rectangle{P1: Point{1, 2}, P2: Point{3, 4}}
	Printf("rectangle: %#v \n", r)

	pixel1 := Pixel{Point{3, 5}, "red"}
	// pixel1.Left = 3
	// pixel1.Right = 5
	// pixel1.Color = "red"
	Println("Embedded struct: ", pixel1)

	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := struct {
		Left  int
		Right int
	}{1, 2}
	p4 := newPoint()
	p4.Left = 1
	p4.Right = 2
	Printf("point#1:%p %#v  point#2:%p %#v  point#3:%p %#v  point#4: %p %#v  point#1==point#2:%v   point#1==point#3:%v   point#1==point#4:%v \n ", &p1, p1, &p2, p2, &p3, p3, p4, p4, p1 == p2, p1 == p3, p1 == *p4)

	Printf("Empty struct (empty object ): %#v\n", struct{}{})
}
