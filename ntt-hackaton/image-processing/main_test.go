package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"testing"
	"time"

	"./protos"
	"google.golang.org/grpc"
)

func startServer() {
	main()
}

func createRandomData() []byte {
	blob := make([]byte, 128*1024) // 128MiB
	rand.Read(blob)
	return blob
}

func sendData(portNumber int, data []byte) {
	conn, err := grpc.Dial(fmt.Sprintf(":%v", portNumber), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	cc := protos.NewChunkerClient(conn)
	client, err := cc.Chunker(context.Background())
	if err != nil {
		panic(err)
	}
	chunk := protos.Chunk{}
	chunk.Chunk = data
	log.Println("sending data to server...  ")
	client.Send(&chunk)
}

func TestSending(t *testing.T) {
	go startServer()
	time.Sleep(1 * time.Second)
	sendData(9999, createRandomData())
	time.Sleep(2 * time.Second)
	sendData(9999, createRandomData())
	time.Sleep(2 * time.Second)
	sendData(9999, createRandomData())
	time.Sleep(1 * time.Second)
}
