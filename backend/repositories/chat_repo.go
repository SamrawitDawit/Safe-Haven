package repositories

import "backend/domain"

type ChatRepo struct {
	messages []domain.ChatMessage
}

func (r *ChatRepo) AddMessage(message domain.ChatMessage) {
	r.messages = append(r.messages, message)
}

func (r *ChatRepo) GetMessages() []domain.ChatMessage {
	return r.messages
}
