package domain

import (
	"time"

	"github.com/google/uuid"
)

type AnonymousUser struct {
	ID                      uuid.UUID `json:"id" bson:"_id,omitempty"`
	AnonymousDifferentiator string    `json:"anonymousDifferentiator" bson:"anonymousDifferentiator"`
	Password                string    `json:"password" bson:"password"`
	AccessToken             string    `json:"accessToken" bson:"accessToken"`
	RefreshToken            string    `json:"refreshToken" bson:"refreshToken"`
	Language                string    `json:"language" bson:"language"`
	Category                string    `json:"category" bson:"category"`
	Active                  bool      `json:"active" bson:"active"`
	Verified                bool      `json:"verified" bson:"verified"`
	PreferedContact         string    `json:"preferedContact" bson:"preferedContact"`
	CounselorAssigned       bool      `json:"counselorAssigned" bson:"counselorAssigned"`
	CounselorID             uuid.UUID `json:"counselorId" bson:"counselorId,omitempty"`
	CreatedAt               time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt" bson:"updatedAt"`
}
