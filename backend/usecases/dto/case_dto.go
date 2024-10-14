package dto

import "github.com/google/uuid"

type CaseDto struct {
	SubmitterID uuid.UUID `json:"submitter_id,omitempty" bson:"submitter_id,omitempty"`
	Title       string    `json:"title,omitempty" bson:"title,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	ImageURL    string    `json:"image_url,omitempty" bson:"image_url,omitempty"`
	VideoURL    string    `json:"video_url,omitempty" bson:"video_url,omitempty"`
	Location    string    `json:"location,omitempty" bson:"location,omitempty"`
}
