package main

import (
	"fmt"
	"./copy"
	"./struct_analyser"
	"./equals"
)

func main(){
	fmt.Println("\n ---------- example of copy data from one struct to another ----")
	copy.ReflectCopy()

	fmt.Println("\n ---------- example print structure code  ----")
	struct_analyser.ReflectInfo(&copy.Movie{"my movie", 100})

	fmt.Println("\n ---------- example of equality  ----")
	equals.Equals()
}