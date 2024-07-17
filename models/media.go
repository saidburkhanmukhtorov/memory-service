package models

import "time"

// CreateMediaModel represents the data structure for creating a new media entry.
type CreateMediaModel struct {
	ID       string    `json:"id" bson:"_id"`
	MemoryID string    `json:"memory_id" bson:"memory_id"`
	Type     string    `json:"type" bson:"type"`
	URL      string    `json:"url" bson:"url"`
	Created  time.Time `json:"created_at" bson:"created_at"`
}

// UpdateMediaModel represents the data structure for updating an existing media entry (PUT).
type UpdateMediaModel struct {
	ID       string    `json:"id" bson:"_id"`
	MemoryID string    `json:"memory_id" bson:"memory_id"`
	Type     string    `json:"type" bson:"type"`
	URL      string    `json:"url" bson:"url"`
	Created  time.Time `json:"created_at" bson:"created_at"`
}

// PatchMediaModel represents the data structure for partially updating an existing media entry (PATCH).
type PatchMediaModel struct {
	ID       string     `json:"id" bson:"_id"`
	MemoryID *string    `json:"memory_id,omitempty" bson:"memory_id,omitempty"`
	Type     *string    `json:"type,omitempty" bson:"type,omitempty"`
	URL      *string    `json:"url,omitempty" bson:"url,omitempty"`
	Created  *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
