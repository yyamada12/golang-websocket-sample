package main

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
)

var (
	origin = "http://localhost:8080/"
	url    = "ws://localhost:8080/ws"
)

// EchoMsg is Sample Websocket Message
type EchoMsg struct {
	Msg string // メッセージ
	ID  int32  // ID
}

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	defer log.Printf("Web Socket Client Sample end.")

	go receiveMsg(ws)

	sendMsg(ws, "Hello", 1)
	sendMsg(ws, "Goodbye", 2)

	time.Sleep(1 * time.Second)
	_ = ws.Close()

}

func sendMsg(ws *websocket.Conn, msg string, id int32) {
	var sndMsg = EchoMsg{msg, id}

	websocket.JSON.Send(ws, sndMsg)
	log.Printf("Send data=%#v\n", sndMsg)
}

func receiveMsg(ws *websocket.Conn) {
	// var rcvMsg EchoMsg
	var msg string
	for {
		// websocket.JSON.Receive(ws, &rcvMsg)
		websocket.Message.Receive(ws, &msg)
		// log.Printf("Receive data=%#v\n", rcvMsg)
		log.Printf("Receive data=%#v\n", msg)
	}
}
