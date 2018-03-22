package main

import (
	"./readget"
	"./server"
)

func main() {
	// read data via GET request
	readget.Example()
	server.Example()
}
