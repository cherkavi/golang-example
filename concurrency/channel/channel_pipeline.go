package channel

import "fmt"

func pipe(input <-chan int, transformer func(int) int) <-chan int {
	out := make(chan int)

	funcPipe := func() {
		for inputValue := range input {
			out <- transformer(inputValue)
		}
		close(out)
	}
	go funcPipe()
	return out
}

func ExampleChannelPipeline() {
	source := integerGenerator(1, 3, 5, 8, 13, 21, 34, 55)

	increaseValue := func(in int) int {
		return in + 1
	}
	doubleValue := func(in int) int {
		return in * 2
	}
	squareValue := func(in int) int {
		return in * in
	}

	for eachValue := range pipe(pipe(pipe(source, increaseValue), doubleValue), squareValue) {
		fmt.Printf("%v  ", eachValue)
	}
	fmt.Println()

}
