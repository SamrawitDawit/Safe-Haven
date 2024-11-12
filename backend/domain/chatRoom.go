package domain

type ChatRoom struct {
	RoomID int `json:"room_id"`
	Participant1 int `json:"participant1"`
	Participant2 int `json:"participant2"`
	CreatedAt string `json:"created_at"`
}