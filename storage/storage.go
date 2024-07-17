package storage

import (
	"context"

	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/models"
)

// StorageI defines the interface for interacting with the storage layer.
type StorageI interface {
	Memory() MemoryI
	Media() MediaI
	Comment() CommentI
}

// MemoryI defines methods for interacting with memory data.
type MemoryI interface {
	CreateMemory(ctx context.Context, memory *models.CreateMemoryModel) (string, error)
	GetMemoryByID(ctx context.Context, id string) (*memory.Memory, error)
	GetAllMemories(ctx context.Context, req *memory.GetAllMemoriesRequest) ([]*memory.Memory, error)
	UpdateMemory(ctx context.Context, memory *models.UpdateMemoryModel) error
	PatchMemory(ctx context.Context, memory *models.PatchMemoryModel) error
	DeleteMemory(ctx context.Context, id string) error
}

// MediaI defines methods for interacting with media data.
type MediaI interface {
	CreateMedia(ctx context.Context, media *models.CreateMediaModel) (string, error)
	GetMediaByID(ctx context.Context, id string) (*memory.Media, error)
	GetAllMedia(ctx context.Context, req *memory.GetAllMediaRequest) ([]*memory.Media, error)
	UpdateMedia(ctx context.Context, media *models.UpdateMediaModel) error
	PatchMedia(ctx context.Context, media *models.PatchMediaModel) error
	DeleteMedia(ctx context.Context, id string) error
}

// CommentI defines methods for interacting with comment data.
type CommentI interface {
	CreateComment(ctx context.Context, comment *models.CreateCommentModel) (string, error)
	GetCommentByID(ctx context.Context, id string) (*memory.Comment, error)
	GetAllComments(ctx context.Context, req *memory.GetAllCommentsRequest) ([]*memory.Comment, error)
	UpdateComment(ctx context.Context, comment *models.UpdateCommentModel) error
	PatchComment(ctx context.Context, comment *models.PatchCommentModel) error
	DeleteComment(ctx context.Context, id string) error
}
