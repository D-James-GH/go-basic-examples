package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrade = websocket.Upgrader{}

func main() {
	http.HandleFunc("/ws", websocketUpgradeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func websocketUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read failed:", err)
			break
		}
		input := string(message)
		err = conn.WriteMessage(messageType, []byte(input))
		if err != nil {
			log.Println(err)
			break
		}
	}
}
