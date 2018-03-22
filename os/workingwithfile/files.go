package WorkingWithFile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func FileOperations() {
	fmt.Println("---- create temp file ---- ")
	var f *os.File
	var err error
	f, err = ioutil.TempFile("", "tempfile")
	if err != nil {
		panic(errors.New("can't create temp file "))
	}
	fmt.Printf(" name of created temp file: %v\n", f.Name())
}
