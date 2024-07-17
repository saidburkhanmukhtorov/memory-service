package models

import "time"

// CreateCommentModel represents the data structure for creating a new comment.
type CreateCommentModel struct {
	ID       string    `json:"id" bson:"_id"`
	MemoryID string    `json:"memory_id" bson:"memory_id"`
	UserID   string    `json:"user_id" bson:"user_id"`
	Content  string    `json:"content" bson:"content"`
	Created  time.Time `json:"created_at" bson:"created_at"`
}

// UpdateCommentModel represents the data structure for updating an existing comment (PUT).
type UpdateCommentModel struct {
	ID       string    `json:"id" bson:"_id"`
	MemoryID string    `json:"memory_id" bson:"memory_id"`
	UserID   string    `json:"user_id" bson:"user_id"`
	Content  string    `json:"content" bson:"content"`
	Created  time.Time `json:"created_at" bson:"created_at"`
}

// PatchCommentModel represents the data structure for partially updating an existing comment (PATCH).
type PatchCommentModel struct {
	ID       string     `json:"id" bson:"_id"`
	MemoryID *string    `json:"memory_id,omitempty" bson:"memory_id,omitempty"`
	UserID   *string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Content  *string    `json:"content,omitempty" bson:"content,omitempty"`
	Created  *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
