package workingwithfile

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	executils "../exec"
)

func createTempFile() *os.File {
	f, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		panic(errors.New("can't create temp file "))
	}
	return f
}

func writeDataIntoFile(path string) {
	file, _ := os.OpenFile(path, os.O_WRONLY, 0666)
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("this is string for file ")
	writer.Flush()
}

func readDataFromFile(path string) {
	// read file via command line
	executils.ExecuteCommandAndOutput("cat", path)

	// read file via utils class
	data, _ := ioutil.ReadFile(path)
	fmt.Println(string(data))

	// read file via buffer
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Scan()
	fmt.Printf("File content: %v \n", scanner.Text())
}

func printStatistic(file *os.File) {
	fileStat, _ := file.Stat()
	fmt.Printf("Name: %v\n", fileStat.Name())
	fmt.Printf("IsDir: %v\n", fileStat.IsDir())
	fmt.Printf("Mode: %v\n", fileStat.Mode())
	fmt.Printf("ModTime: %v\n", fileStat.ModTime())
	fmt.Printf("Size: %v\n", fileStat.Size())
	fmt.Printf("Sys: %v\n", fileStat.Sys())
}

func FileOperations() {
	fmt.Println("---- Operations with file  ---- ")
	var file *os.File = createTempFile()
	fmt.Println("> create Temp file: ", file.Name())

	fmt.Println("> statistic  ")
	printStatistic(file)

	fmt.Println("> write data into file ")
	writeDataIntoFile(file.Name())

	fmt.Println("> read data from file  ")
	readDataFromFile(file.Name())

	// os.Rename - the same like OS.mv
	// os.Remove - the same like OS.rm
}
