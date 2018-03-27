package main

import (
	"time"

	"./fileserver"
	"./readget"
	"./restapp"
	"./server"
)

func main() {
	// read data via GET request
	readget.Example()
	server.Example()
	restapp.Example()
	fileserver.StartFileServer()

	time.Sleep(time.Second * 20)
}
