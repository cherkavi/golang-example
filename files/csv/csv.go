package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeDataIntoFile(fileName string) {
	file, _ := os.OpenFile(fileName, os.O_WRONLY, 666)
	defer file.Close()
	writer := csv.NewWriter(file)
	// writer.Comma = ';'
	writer.Write([]string{"h1", "h2", "h3", "h4", "h5"})
	writer.Write([]string{"1", "2", "3,4", "3,5", "4"})
	writer.Write([]string{"one", "two", "three", "four", "five"})
	writer.Flush()
}

func readDataFromFile(fileName string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		// csvError = err.(*csv.ParseError)
		// csvError.Column, csvError.Line, csvError.Err
		// csv.ErrFieldCount
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}

func Example() {
	file, _ := ioutil.TempFile("", "temp file")
	// os.Create
	defer func() {
		_ = os.Remove(file.Name())
	}()

	fmt.Println("> write data into file ")
	writeDataIntoFile(file.Name())

	fmt.Println("> file content ")
	text, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(text))

	fmt.Println("> read data from file")
	readDataFromFile(file.Name())
}
