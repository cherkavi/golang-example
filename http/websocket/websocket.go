package websocket

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// Message is exchange object between client and server
type Message struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func handleClient(connection *websocket.Conn) {
	fmt.Println("server: new client connected: ", connection)
	var message Message
MainLoop:
	for {
		// websocket.Message.Receive
		switch err := websocket.JSON.Receive(connection, &message); err {
		case nil:
			if err = websocket.JSON.Send(connection, message); err != nil {
				fmt.Println("server: sending message error: ", err)
			} else {
				fmt.Println("server: message was sent back ")
			}

		case io.EOF:
			fmt.Println("server: remote client disconnected")
			connection.Close()
			break MainLoop // will break for loop
		default:
			fmt.Println("server: remote client error: ", err)
			connection.Close()
			break MainLoop // will break for loop
		}
	}

}

// ExampleOfWebSocket - start websocket server
func ExampleOfWebSocket() {
	fmt.Println("---- start WebSocket server ----")
	webSocketServer := http.NewServeMux()
	webSocketServer.Handle("/", websocket.Handler(handleClient))
	go http.ListenAndServe(":8084", webSocketServer)

	<-time.After(1 * time.Second)
	fmt.Println("> create websocket client connection")
	wsClient, err := websocket.Dial("ws://127.0.0.1:8084", "", "/") // wss prefix is not working
	if err != nil {
		fmt.Printf(" can't open web socket: %v \n", err)
		return
	}

	fmt.Println("> client send message ")
	var message = Message{"10", "new client"}
	// websocket.Message.Send
	if err = websocket.JSON.Send(wsClient, &message); err != nil {
		fmt.Println("client did not send data to remote websocket error: ", err)
	}
	fmt.Println("> client sent message")

	if err = websocket.JSON.Receive(wsClient, &message); err != nil {
		fmt.Println("client reading data from socket error: ", err)
	}
	fmt.Println("echo message from server: ", message)
	fmt.Println("client is closing connection ")
	wsClient.Close()
}
