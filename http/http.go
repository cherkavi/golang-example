package main

import (
	"time"

	"./readget"
	"./restapp"
	"./server"
)

func main() {
	// read data via GET request
	readget.Example()
	server.Example()
	restapp.Example()

	time.Sleep(time.Second * 20)
}
