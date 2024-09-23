package domain

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	ID                uuid.UUID `json:"id" bson:"_id"`
	ReporterID        uuid.UUID `json:"reporter_id" bson:"reporter_id,omitempty"`
	Title             string    `json:"title" bson:"title"`
	Description       string    `json:"description" bson:"description"`
	ImageURL          string    `json:"image_url" bson:"image_url,omitempty"`
	Location          string    `json:"location" bson:"location,omitempty"`
	Status            string    `json:"status" bson:"status"`
	ReportedAt        time.Time `json:"reported_at" bson:"reported_at"`
	UpdatedAt         time.Time `json:"updated_at" bson:"updated_at"`
	CounselorAssigned bool      `json:"counselor_assigned" bson:"counselor_assigned"`
	CounselorID       uuid.UUID `json:"counselor_id" bson:"counselor_id,omitempty"`
}
