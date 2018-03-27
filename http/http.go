package main

import (
	"time"

	"./fileserver"
	"./readget"
	"./restapp"
	"./server"
	"./sslserver"
	"./websocket"
)

func main() {
	// read data via GET request
	readget.Example()
	server.Example()
	restapp.Example()
	fileserver.StartFileServer()
	websocket.ExampleOfWebSocket()
	sslserver.ExampleOfSSL()
	time.Sleep(time.Second * 20)
}
