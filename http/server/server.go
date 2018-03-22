package server

import (
	"fmt"
	"log"
	"net/http"
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "time: %v", time.Now().Format(time.RFC1123))
}

func Example() {
	var portNumber = 8001
	fmt.Printf("---- start local server on port :%v \n", portNumber)

	http.HandleFunc("/", rootHandler)
	go http.ListenAndServe(fmt.Sprintf(":%v", portNumber), nil)

	fmt.Printf("waiting for start of the server: 127.0.0.1:%v \n", portNumber)
	time.Sleep(1 * time.Second)

	fmt.Println("open browser")
	openbrowser(fmt.Sprintf("http://127.0.0.1:%v/", portNumber))

	fmt.Println("sleep for 60 seconds")
	time.Sleep(60 * time.Second)
}
