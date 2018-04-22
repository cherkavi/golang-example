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

func writeAnswerIntoTempFile(values []byte) (*os.File, error) {
    var tempFile *os.File
    var err error
    if tempFile, err = createTempFile(); err != nil {
	log.Panic("can't create temp file ")
	return nil, err
    }
    defer tempFile.Close()

    tempFile.Write(values)
    tempFile.Sync()
    return tempFile, err
}
