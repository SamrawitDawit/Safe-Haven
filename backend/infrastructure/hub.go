// infrastructure/hub.go
package infrastructure

import (
	"backend/domain"
)

type HubImpl struct {
	clients    map[*domain.Client]bool // Registered clients
	broadcast  chan []byte             // Channel for broadcasting messages
	register   chan *domain.Client     // Channel for registering clients
	unregister chan *domain.Client     // Channel for unregistering clients
}

// Ensure HubImpl satisfies the domain.Hub interface
var _ domain.Hub = &HubImpl{}

// NewHub creates a new instance of HubImpl
func NewHub() *HubImpl {
	return &HubImpl{
		clients:    make(map[*domain.Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *domain.Client),
		unregister: make(chan *domain.Client),
	}
}

// Run listens for register, unregister, and broadcast actions
func (h *HubImpl) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// RegisterClient registers a client
func (h *HubImpl) RegisterClient(client *domain.Client) {
	h.register <- client
}

// UnregisterClient unregisters a client
func (h *HubImpl) UnregisterClient(client *domain.Client) {
	h.unregister <- client
}

// BroadcastMessage sends a message to all clients
func (h *HubImpl) BroadcastMessage(message []byte) {
	h.broadcast <- message
}
