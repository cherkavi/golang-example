package udp

import (
	"fmt"
	"net"
	"time"
)

func panicWhenError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error: %v\n", err))
	}
}

func startUdpListener(portNumber int) {
	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", portNumber))
	panicWhenError(err)

	fmt.Println("start UDP listener ")
	serverConnection, err := net.ListenUDP("udp", address)
	panicWhenError(err)
	defer serverConnection.Close()

	buffer := make([]byte, 8196)

	for {
		amountOfBytes, remoteAddress, err := serverConnection.ReadFromUDP(buffer)
		fmt.Println("Received: ", string(buffer[0:amountOfBytes]), " from ", remoteAddress)

		panicWhenError(err)
	}
}

func writeUdpPackage(host string, portNumber int, data []byte) {
	remoteAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", host, portNumber))
	panicWhenError(err)

	localAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	panicWhenError(err)

	fmt.Printf(">>> send UDP package to remote address: %v\n", remoteAddress)
	connection, err := net.DialUDP("udp", localAddress, remoteAddress)
	panicWhenError(err)
	defer connection.Close()

	_, err = connection.Write(data)
	panicWhenError(err)
}

func Example() {
	go startUdpListener(9000)
	for i := 0; i < 3; i++ {
		writeUdpPackage("127.0.0.1", 9000, []byte(time.Now().Format(time.RFC1123)))
	}
	time.Sleep(time.Second * 10)
}
