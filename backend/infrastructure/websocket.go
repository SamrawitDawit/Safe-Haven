// infrastructure/websocket.go
package infrastructure

import (
	"backend/delivery/controllers"
	"backend/usecases"
)

func SetupWebSocketServer() *controllers.WebSocketController {
	hub := NewHub()  
	go hub.Run()   

	u :=  &usecases.ChatUsecase{
		Hub: hub,  
	}
	return &controllers.WebSocketController{
		Usecase: u,  
	}
}
