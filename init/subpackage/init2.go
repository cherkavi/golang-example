package subpackage

import (
	"fmt"
)

func init() {
	fmt.Println("  init method into package subinit")
}

func init() {
	fmt.Println("  init method (second) into package subinit")
}

func init() {
	fmt.Println("  init method (third) into package subinit")
}

func JustEmptyFunction() {
}
