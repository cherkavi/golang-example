package inner3

import "fmt"

func Fibonacci() func() int {
	var beforePrevious, previous int = -1, 0
	return func() int{
		if beforePrevious==(-1){
			beforePrevious=0
			return 0
		}
		if previous==(0){
			previous=1
			return 1
		}
		var returnValue int = beforePrevious+previous
		beforePrevious, previous = previous, returnValue
		return returnValue
	}
}

func init() {
	fmt.Println("func init from package 'inner3' ")
}
