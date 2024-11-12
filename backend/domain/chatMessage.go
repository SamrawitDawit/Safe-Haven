package domain

type ChatMessage struct {
	ID        int    `json:"id"`
	Content  string `json:"content"`
	SenderID int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Seen 	bool   `json:"seen"`
}