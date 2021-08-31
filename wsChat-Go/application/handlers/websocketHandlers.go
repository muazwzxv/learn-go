package handlers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// WebsocketResponse Response type of websocket
type WebsocketResponse struct {
	Action string `json:"action"`
	Message string `json:"message"`
	MessageType string `json:"message_type"`
}

var upgradeConnection = websocket.Upgrader{
ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebsocketEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Upgrading to websocket ......")

	err = ws.WriteJSON(WebsocketResponse{"","<em><small>Connected to server</small></em", "" })
	if err != nil {
		log.Println(err)
	}
}
