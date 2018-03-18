package main

import (
	"fmt"
)

func printType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float32:
		return "float"
	case float64:
		return "float"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func main() {
	var x int = 10

	if controlValue := x + 5; controlValue > 0 { // variable 'controlValue' visible into inner scope only
		fmt.Println("example of calculation before evaluation ")
	}
	fmt.Println()

	fmt.Println("---- switch fallthrough feature ---- ")
	switch controlValue := 10; controlValue { // variable 'controlValue' visible into inner scope only
	case 9, 10, 11, 12:
		fmt.Println("case when controlValue == [9..12]")
	case 0:
		fmt.Println("case when controlValue == 0")
	default:
		fmt.Println("other cases ")
	}
	fmt.Println()

	switch controlValue := x - 10; { // variable 'controlValue' visible into inner scope only
	case controlValue > 0:
		fmt.Println("case when controlValue biggest than 0")
	case controlValue == 0:
		fmt.Println("case when controlValue equals 0")
	default:
		fmt.Println("other cases ")
	}
	fmt.Println()

	switch controlValue := 21; { // variable 'controlValue' visible into inner scope only
	default:
		fmt.Println("other cases ")
		fallthrough
	case controlValue > 30:
		fmt.Println("controlValue > 30")
		fallthrough
	case controlValue > 20:
		fmt.Println("controlValue > 20")
		fallthrough
	case controlValue > 10:
		fmt.Println("controlValue > 10")
		fallthrough
	case controlValue > 40:
		fmt.Println("strange, but condition was not considered : controlValue > 40, after this line will stop ( no fallthrough )")
	case controlValue > 0:
		fmt.Println("controlValue > 0")
	}

	fmt.Printf("print Type: %v  \n", printType(10))

	fmt.Printf("\n---- FOR loop ---- \n")
	for i := 1; i < 5; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	for index, value := range []int{5, 7, 9, 13, 17} {
		fmt.Printf("index: %v,  value: %v\n", index, value)
	}
	fmt.Println()

	fmt.Println("---- example of infinite loop with GOTO operator ")
	i := 10
	for {
		fmt.Printf("%v ", i)
		goto EndLabel
	}
EndLabel:
	fmt.Println()

}
