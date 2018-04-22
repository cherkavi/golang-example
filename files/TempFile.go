package main

import (
    "io/ioutil"
    "os"
)


func createTempFile() (*os.File, error) {
    tempFile, err := ioutil.TempFile("", "jpg")
    if err != nil {
	return nil, err
    }
    defer os.Remove(tempFile.Name())
    return os.OpenFile(tempFile.Name()+".jpg", os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
}

