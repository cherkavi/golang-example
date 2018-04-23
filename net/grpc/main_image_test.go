package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"io/ioutil"

	"./protos"
	"google.golang.org/grpc"
)

func sendByteArray(portNumber int, data []byte) {
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

func TestImageRecognition(t *testing.T) {
	go startServer()
	time.Sleep(1 * time.Second)
	var file, _ = os.OpenFile("./recognizer/test-image.jpg", os.O_RDONLY, 400)
	data, _ := ioutil.ReadAll(file)
	sendByteArray(9999, data)
	time.Sleep(4 * time.Second)
}
