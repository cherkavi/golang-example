package inner

import "fmt"

func notVisibleFunction() {
	fmt.Println("without parameters, without return value")
}

func ExternalAdd(values ... int) (returnValue int){
	for _, value := range values  {
		returnValue += value
	}
	return
}