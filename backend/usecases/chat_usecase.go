package usecases

import "backend/domain"

type ChatUsecase struct {
	Hub domain.Hub
}
func (u *ChatUsecase) RegisterClient(client *domain.Client) {
	u.Hub.RegisterClient(client)
}
func (u *ChatUsecase) UnregisterClient(client *domain.Client) {
	u.Hub.UnregisterClient(client)
}
func (u *ChatUsecase) BroadcastMessage(message []byte) {
	u.Hub.BroadcastMessage(message)
}

