package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	dl := websocket.Dialer{}
	conn, h, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(h)
	conn.WriteMessage(websocket.TextMessage, []byte("我来了"))
}
