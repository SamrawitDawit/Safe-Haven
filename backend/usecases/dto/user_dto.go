package dto

type RegisterDTO struct {
	FullName    string `json:"fullName" binding:"required"`
	Email       string `json:"email omitempty"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
}
type LoginDTO struct {
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password" binding:"required"`
}

type AnonymousUser struct {
	AnonymousDifferenitator string `json:"anonymousDifferenitator" binding:"required"`
	Password                string `json:"password" binding:"required"`
}

type TokenResponseDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
