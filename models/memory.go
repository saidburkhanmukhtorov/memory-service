package models

import "time"

// CreateMemoryModel represents the data structure for creating a new memory.
type CreateMemoryModel struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Tags        []string  `json:"tags" bson:"tags"`
	Latitude    float64   `json:"latitude" bson:"latitude"`
	Longitude   float64   `json:"longitude" bson:"longitude"`
	PlaceName   string    `json:"place_name" bson:"place_name"`
	Privacy     string    `json:"privacy" bson:"privacy"`
}

// UpdateMemoryModel represents the data structure for updating an existing memory (PUT).
type UpdateMemoryModel struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Date        time.Time `json:"date" bson:"date"`
	Tags        []string  `json:"tags" bson:"tags"`
	Latitude    float64   `json:"latitude" bson:"latitude"`
	Longitude   float64   `json:"longitude" bson:"longitude"`
	PlaceName   string    `json:"place_name" bson:"place_name"`
	Privacy     string    `json:"privacy" bson:"privacy"`
}

// PatchMemoryModel represents the data structure for partially updating an existing memory (PATCH).
type PatchMemoryModel struct {
	ID          string     `json:"id" bson:"_id"`
	Title       *string    `json:"title,omitempty" bson:"title,omitempty"`
	Description *string    `json:"description,omitempty" bson:"description,omitempty"`
	Date        *time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Tags        *[]string  `json:"tags,omitempty" bson:"tags,omitempty"`
	Latitude    *float64   `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude   *float64   `json:"longitude,omitempty" bson:"longitude,omitempty"`
	PlaceName   *string    `json:"place_name,omitempty" bson:"place_name,omitempty"`
	Privacy     *string    `json:"privacy,omitempty" bson:"privacy,omitempty"`
}
