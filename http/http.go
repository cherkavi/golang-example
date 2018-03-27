package main

import (
	"time"

	"./fileserver"
	"./readget"
	"./restapp"
	"./server"
	"./websocket"
)

func main() {
	// read data via GET request
	readget.Example()
	server.Example()
	restapp.Example()
	fileserver.StartFileServer()
	websocket.ExampleOfWebSocket()
	time.Sleep(time.Second * 20)
}
