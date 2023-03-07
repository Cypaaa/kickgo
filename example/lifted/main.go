package main

import (
	"net/http"

	"github.com/Cypaaa/kickgo"

	"github.com/gorilla/websocket"
)

var SUBGAP uint64 = 99999999999
var SUBCOUNT uint64 = 0
var WSCONN []*websocket.Conn = make([]*websocket.Conn, 0)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	var k *kickgo.Kick = kickgo.New()
	// k.AddChatroom("Cypaa")
	k.AddChatroomId("513716")
	k.AddHandler(subscribed)
	k.AddHandler(messageSent)

	// Edits the SUBCOUNT value
	go k.Listen("YOUR TOKEN HERE")

	// Sends the new value of SUBCOUNT to the client
	go SendLoop()

	// Creates ws conns and store them in WSCONN
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { wsHandler(k, w, r) })
	http.ListenAndServe(":9898", nil)
}
