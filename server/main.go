package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func webSocketHandler(ws *websocket.Conn) {
	defer ws.Close()

	// 初回メッセージを送信
	err := websocket.Message.Send(ws, "Server: Hello, Client!")
	if err != nil {
		log.Println(err)
	}

	for {
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(msg)

		err = websocket.Message.Send(ws,
			fmt.Sprintf("Server: '%s' received.", msg))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(webSocketHandler))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
