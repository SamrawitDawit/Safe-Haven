package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `json:"id" bson:"_id"`
	FullName          string    `json:"fullName" bson:"fullName" binding:"required"`
	Email             string    `json:"email" bson:"email,omitempty"`
	Password          string    `json:"password" bson:"password" binding:"required"`
	PhoneNumber       string    `json:"phoneNumber" bson:"phoneNumber,omitempty"`
	Bio               string    `json:"bio" bson:"bio,omitempty"`
	ImageURL          string    `json:"imageUrl" bson:"imageUrl,omitempty"`
	Role              string    `json:"role" bson:"role" binding:"required"`
	Category          string    `json:"category" bson:"category"`
	Active            bool      `json:"active" bson:"active"`
	Verified          bool      `json:"verified" bson:"verified"`
	CounselorAssigned bool      `json:"counselorAssigned" bson:"counselorAssigned" binding:"required"`
	PreferedContact   string    `json:"preferedContact" bson:"preferedContact,omitempty"`
	CounselorID       uuid.UUID `json:"counselorId" bson:"counselorId,omitempty"`
	Language          string    `json:"language" bson:"language" binding:"required"`
	AccessToken       string    `json:"accessToken" bson:"accessToken,omitempty"`
	RefreshToken      string    `json:"refreshToken" bson:"refreshToken,omitempty"`
	TwoFactorEnabled  bool      `json:"twoFactorEnabled" bson:"twoFactorEnabled,omitempty"`
	ResetToken        string    `json:"resetToken" bson:"resetToken,omitempty"`
	ResetTokenExpiry  time.Time `json:"resetTokenExpiry" bson:"resetTokenExpiry,omitempty"`
	ResetCode         string    `json:"resetCode" bson:"resetCode,omitempty"`
	GoogleSignin      bool      `json:"googleSignin" bson:"googleSignin,omitempty"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt" binding:"required"`
	UpdatedAt         time.Time `json:"updatedAt" bson:"updatedAt" binding:"required"`
	Lock              bool      `json:"lock" bson:"lock"`
}
