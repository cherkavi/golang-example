package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("---- File Walk, path walk  ")
	filepath.Walk("/tmp", func(path string, info os.FileInfo, err error) error {
		file, _ := os.Open(path)
		if data, err := ioutil.ReadAll(file); err != nil {
			fmt.Printf("%v -> md5: %v  -> size: %v\n", path, md5.Sum(data), info.Size())
		}
		return nil
	})
}
