package main

import (
	"time"

	"./tcp"
	"./udp"
)

func main() {
	tcp.Example()
	time.Sleep(time.Second * 3)
	udp.Example()
}
