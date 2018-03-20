package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

func httpRawRequest(remoteUrl string) {
	var connection net.Conn
	var err error
	if connection, err = net.Dial("tcp", remoteUrl); err != nil {
		panic(errors.New("can't open TCP connection " + remoteUrl))
	}
	defer connection.Close()

	fmt.Println("    write data into socket >>> ")
	fmt.Fprintf(connection, "GET / HTTP/1.0\r\n\r\n")

	var firstLine string
	if firstLine, err = bufio.NewReader(connection).ReadString('\n'); err != nil {
		panic(errors.New("can't read data from socket "))
	}
	fmt.Printf("    read data from socket: %v \n", firstLine)
}

func startListener(portNumber int) {
	var err error
	var listener net.Listener
	fmt.Println("   start listener ")
	if listener, err = net.Listen("tcp", fmt.Sprintf(":%v", portNumber)); err != nil {
		panic(errors.New(fmt.Sprintf("can't listen port: %v ", portNumber)))
	}

	for {
		if connection, err := listener.Accept(); err == nil {
			fmt.Printf("    remote client %#v connected to %#v  \n", connection.LocalAddr(), connection.RemoteAddr())
			go timeResponse(connection)
		} else {
			panic(errors.New("can't open connection with remote client "))
		}
	}
}

func timeResponse(connection net.Conn) {
	defer connection.Close()
	var data string
	fmt.Fscan(connection, &data)
	fmt.Printf("data from remote host <<< : %v \n", data)
	fmt.Fprintf(connection, "data from remote host: %v\n", time.Now().Format(time.RFC1123))
}

func main() {
	httpRawRequest("google.com:80")

	fmt.Println("---- client-server example ----")
	portNumber := 8081
	go startListener(portNumber)
	time.Sleep(time.Second * 1)
	for i := 0; i < 10; i++ {
		httpRawRequest(fmt.Sprintf("127.0.0.1:%v", portNumber))
	}
}
