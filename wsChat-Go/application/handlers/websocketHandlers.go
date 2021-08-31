package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var wsChan = make(chan WebsocketPayload)
var clients = make(map[WebSocketConnection]string)

type WebSocketConnection struct {
	*websocket.Conn
}

// WebsocketResponse Response type of websocket
type WebsocketResponse struct {
	Action string `json:"action"`
	Message string `json:"message"`
	MessageType string `json:"message_type"`
}

// WebsocketPayload being the payload received by the client
type WebsocketPayload struct {
	Action string `json:"action"`
	Message string `json:"message"`
	Username string `json:"username"`
	Conn WebSocketConnection `json:"-"`
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

	conn := WebSocketConnection{ws}
	clients[conn] = ""

	err = ws.WriteJSON(WebsocketResponse{"","<em><small>Connected to server</small></em", "" })
	if err != nil {
		log.Println(err)
	}

	go ListenWebsocket(&conn)
}

// ListenWebsocket recovers if goroutine panic
func ListenWebsocket(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recored")
			log.Println("Error was", fmt.Sprintf("%v", r))
		}
	}()

	var payload WebsocketPayload

	// Infinite loop to listen to upcoming message and sends it to
	// other goroutine through channel
	for {
		if err := conn.ReadJSON(&payload); err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func ListenToSocketChannel() {
	var res WebsocketResponse

	for {
		e := <- wsChan
		res.Action = "Got here"
		res.Message = fmt.Sprintf("Some message, and action was %s", e.Action)
		broadcastMessage(res)
	}
}

func broadcastMessage(res WebsocketResponse) {
	for client := range clients	 {
		err := client.WriteJSON(res)
		if err != nil {
			log.Println("Websocket error")
			_ = client.Close()
			delete(clients, client)
		}
	}
}
