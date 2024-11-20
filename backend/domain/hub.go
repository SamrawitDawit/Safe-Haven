package domain

type Hub interface {
	RegisterClient(client *Client)
	UnregisterClient(client *Client)
	BroadcastMessage(message []byte)
}

