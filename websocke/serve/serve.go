package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func hand(req http.ResponseWriter, res *http.Request) {
	coon, err := up.Upgrade(req, res, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		message, i, err := coon.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(message, string(i))
	}
	defer coon.Close()
	log.Println("服务被关闭")
}

func main() {
	http.HandleFunc("/test", hand)
	http.ListenAndServe(":8888", nil)
}
