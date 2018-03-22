package server

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"../../os/exec"
)

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("request url: ", request.RequestURI)
	urlValues, _ := url.Parse(request.RequestURI)
	fmt.Printf("request param1: %v\n", urlValues.Query().Get("param1"))
	fmt.Printf("request param2: %v\n", urlValues.Query().Get("param2"))
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
	exec.Openbrowser(fmt.Sprintf("http://127.0.0.1:%v/get?param1='hello'&param2=10", portNumber))

}
