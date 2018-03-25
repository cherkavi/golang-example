package unittesting

import (
	"fmt"
	"sync"
	"testing"
)

// to execute functions once between tests
var once sync.Once

func TestAddingSkipped(t *testing.T) {
	if testing.Short() { // check for short testing
		fmt.Println("short testing") // flag "-short" present into command line
		t.Skip()
	} else {
		fmt.Println("regular testing ")
	}
	t.Log("this line will appear only when error happend ")
	if add(2, 3) != 4 {
		t.Errorf("expected 4 but [%v] ", add(2, 3))
		// t.Fatal("test stopped")
	}
}

func ExampleCheckingOutput() {
	fmt.Println(" this is my output ")
	// Output:
	//  this is my output
}

func ExampleCheckingUnorderedOutput() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	// Unordered output:
	// 3
	// 2
	// 1
}
