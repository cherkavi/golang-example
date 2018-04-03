package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func fileWriter(fileName string) *os.File {

	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//defer to close when you're done with it, not because you think it's idiomatic!
	// defer logFile.Close()
	return logFile
}

func simpleLogExample() {
	fmt.Println("----- init simple log ----- ")

	//create your file with desired read/write permissions
	file, _ := ioutil.TempFile("", "tempfile")
	defer func() {
		os.Remove(file.Name())
	}()

	//set output of logs to f
	log.SetOutput(fileWriter(file.Name()))
	//test case
	log.Println("log message into file ")

	text, _ := ioutil.ReadFile(file.Name())
	fmt.Println(string(text))

}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func prodLogInit() {
	initLog(ioutil.Discard, // fileWriter(nameOfFile)
		os.Stdout, // io.MultiWriter(fileWriter(nameOfFile), os.Stdout)
		os.Stdout,
		os.Stderr)

	Trace.Println("trace log example ") // discard by output
	Info.Println("info log example ")
	Warning.Println("warning log example ")
	Error.Println("error log example ")
}

func main() {
	simpleLogExample()
	prodLogInit()
}
