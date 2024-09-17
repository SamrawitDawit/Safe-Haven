package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                      uuid.UUID `json:"id" bson:"_id,omitempty"`
	FullName                string    `json:"fullName" bson:"fullName"`
	Email                   string    `json:"email" bson:"email"`
	Password                string    `json:"password" bson:"password"`
	PhoneNumber             string    `json:"phoneNumber" bson:"phoneNumber"`
	Bio                     string    `json:"bio" bson:"bio"`
	ImageURL                string    `json:"imageUrl" bson:"imageUrl"`
	Role                    string    `json:"role" bson:"role"`
	Category                string    `json:"category" bson:"category"`
	Active                  bool      `json:"active" bson:"active"`
	Verified                bool      `json:"verified" bson:"verified"`
	CounselorAssigned       bool      `json:"counselorAssigned" bson:"counselorAssigned"`
	PreferedContact         string    `json:"preferedContact" bson:"preferedContact"`
	CounselorID             uuid.UUID `json:"counselorId" bson:"counselorId,omitempty"`
	Language                string    `json:"language" bson:"language"`
	AccessToken             string    `json:"accessToken" bson:"accessToken"`
	RefreshToken            string    `json:"refreshToken" bson:"refreshToken"`
	TwoFactorEnabled        bool      `json:"twoFactorEnabled" bson:"twoFactorEnabled"`
	ResetToken              string    `json:"resetToken" bson:"resetToken"`
	ResetCode               int64     `json:"resetCode" bson:"resetCode"`
	GoogleSignin            bool      `json:"googleSignin" bson:"googleSignin"`
	CreatedAt               time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt" bson:"updatedAt"`
}
