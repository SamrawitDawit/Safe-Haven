package controllers

import (
	"backend/domain"
	"backend/usecases"
	"net/http"

	"github.com/gorilla/websocket"
)


type WebSocketController struct {
	Usecase *usecases.ChatUsecase 
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *WebSocketController) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "could not open websocket connection", http.StatusInternalServerError)
		return
	}

	client := &domain.Client{
		ID : r.RemoteAddr,
		Conn: conn,
		Send: make(chan []byte),
	}

	c.Usecase.RegisterClient(client)
	go c.readMessages(client)
	go c.writeMessages(client)
}

func (c *WebSocketController) readMessages(client *domain.Client) {
	defer c.Usecase.UnregisterClient(client)

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
		c.Usecase.BroadcastMessage(msg)
	}
}

func (c *WebSocketController) writeMessages(client *domain.Client) {
	defer client.Conn.Close()

	for message := range client.Send {
		client.Conn.WriteMessage(websocket.TextMessage, message)
	}
}