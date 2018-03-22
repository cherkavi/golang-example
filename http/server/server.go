package server

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("request url: ", request.RequestURI)
	urlValues, _ := url.Parse(request.RequestURI)
	fmt.Printf("request param1: %v", urlValues.Query().Get("param1"))
	fmt.Printf("request param2: %v", urlValues.Query().Get("param2"))
	fmt.Fprintf(response, "time: %v", time.Now().Format(time.RFC1123))
}

type ServerHandler struct {
}

func (h ServerHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(407)
	var writer = bufio.NewWriter(response)
	writer.WriteString("example of error message ")
	writer.Flush()
}

func Example() {
	var portNumber = 8001
	fmt.Printf("---- start local server on port :%v \n", portNumber)

	// go http.ListenAndServe(fmt.Sprintf(":%v", portNumber), ServerHandler{})
	http.HandleFunc("/get", rootHandler)
	go http.ListenAndServe(fmt.Sprintf(":%v", portNumber), nil)

	fmt.Printf("waiting for start of the server: 127.0.0.1:%v \n", portNumber)
	time.Sleep(1 * time.Second)

	fmt.Println("open browser")
	openbrowser(fmt.Sprintf("http://127.0.0.1:%v/get?param1='hello'&param2=10", portNumber))

	fmt.Println("sleep for 10 seconds")
	time.Sleep(10 * time.Second)
}
