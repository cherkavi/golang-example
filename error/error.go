package main

import (
	"errors"
	"fmt"
	"reflect"
)

func returnError() error {
	return errors.New("hard-coded error")
}

func throwCatch() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error happend: %#v   %v", err.(error).Error(), reflect.TypeOf(err))
		}
	}()

	panic(errors.New("intenionally throw exception"))
	// panic("intenionally throw exception")
}

func main() {
	if er := returnError(); er != nil {
		fmt.Printf("Error: %v\n", er.Error())
	}

	throwCatch()
}
