package dto

type RegisterDTO struct {
	FullName    string `json:"fullName" binding:"required"`
	Email       string `json:"email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Language    string `json:"language" binding:"required"`
	Category    string `json:"category" binding:"required"`
}
type LoginDTO struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password" binding:"required"`
}
