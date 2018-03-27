package fileserver

import (
	"fmt"
	"net/http"
	"time"

	"../../os/exec"
)

// StartFileServer - start file server to return content of the file system
func StartFileServer() {
	var portNumber = 9099
	var path = "/tmp"

	fmt.Println("---- start File server on port: ", portNumber)
	// start another http server
	// start second server
	fileServer := http.NewServeMux()
	fileServer.Handle("/", http.FileServer(http.Dir(path)))

	go http.ListenAndServe(fmt.Sprintf(":%v", portNumber), fileServer)

	<-time.After(1 * time.Second)
	url := fmt.Sprintf("http://127.0.0.1:%v", portNumber)
	fmt.Println("> open browser with file system content: ", url)
	exec.Openbrowser(url)
}
