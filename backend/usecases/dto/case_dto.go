package dto

import "github.com/google/uuid"

type CreateCaseDto struct {
	SubmitterID uuid.UUID `json:"submitter_id" bson:"submitter_id,omitempty"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	ImageURL    string    `json:"image_url" bson:"image_url,omitempty"`
	Location    string    `json:"location" bson:"location,omitempty"`
}

type UpdateCaseDto struct {
	ID          uuid.UUID `json:"id" bson:"_id" binding:"required"`
	SubmitterID uuid.UUID `json:"submitter_id" bson:"submitter_id,omitempty"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	ImageURL    string    `json:"image_url" bson:"image_url,omitempty"`
	Location    string    `json:"location" bson:"location,omitempty"`
}
