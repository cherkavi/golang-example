package std

import (
	"bufio"
	"fmt"
	"os"
)

func ExampleOfReadingText() {
	fmt.Println("> read data from console - numbers with space separated ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println("text from console: ", scanner.Text())
	}
}
