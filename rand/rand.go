package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("---- init random generator ----\n")
	rand.Seed(time.Now().UnixNano())
	for index := 0; index < 10; index++ {
		fmt.Printf("random value: %v \n", rand.Int31n(10)) // rand.Int()
	}
}
