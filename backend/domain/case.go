package domain

import (
	"time"

	"github.com/google/uuid"
)

type Case struct {
	ID                uuid.UUID `json:"id" bson:"_id"`
	SubmitterID       uuid.UUID `json:"submitter_id" bson:"submitter_id,omitempty"`
	Title             string    `json:"title" bson:"title"`
	Description       string    `json:"description" bson:"description"`
	ImageURL          string    `json:"image_url" bson:"image_url,omitempty"`
	VideoURL          string    `json:"video_url" bson:"video_url,omitempty"`
	Location          string    `json:"location" bson:"location,omitempty"`
	Status            string    `json:"status" bson:"status"`
	SubmittedAt       time.Time `json:"submitted_at" bson:"submitted_at"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
	CounselorAssigned bool      `json:"counselor_assigned" bson:"counselor_assigned"`
	CounselorID       uuid.UUID `json:"counselor_id" bson:"counselor_id,omitempty"`
}
