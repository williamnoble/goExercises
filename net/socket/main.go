package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func reader(conn *websocket.Conn) {
	for {
		// Read Message from the Socket (i.e. from the client)
		messageType, p, _ := conn.ReadMessage()

		// Print the Read Message to Stdout
		fmt.Printf("MESSAGETYPE: %d\n ", messageType)
		fmt.Println("P: " + string(p))

		// Write a message to the client!
		if err := conn.WriteMessage(messageType, []byte("Hello from the Server :)")); err != nil {
			return
		}

	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	reader(ws)
}

func main() {
	http.HandleFunc("/", ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
