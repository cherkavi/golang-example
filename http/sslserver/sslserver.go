package sslserver

import (
	"fmt"
	"net/http"
	"time"

	"../../os/exec"
)

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("request url: ", request.RequestURI)
	fmt.Fprintf(response, "time: %v", time.Now().Format(time.RFC1123))
}

// openssl.org, create certificate:
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
// Common name - name of the host
// output: cert.pem, key.pem
func ExampleOfSSL() {
	portNumber := 8043
	fmt.Println("---- start SSL server ----")
	server := http.NewServeMux()
	server.HandleFunc("/", rootHandler)
	go http.ListenAndServeTLS(fmt.Sprintf(":%v", portNumber), "cert.pem", "key.pem", server)

	<-time.After(1 * time.Second)
	fmt.Println("open browser with SSL site: ")
	url := "https://127.0.0.1:8043"
	fmt.Println("open browser with SSL site: ", url)
	exec.Openbrowser(url)
}
