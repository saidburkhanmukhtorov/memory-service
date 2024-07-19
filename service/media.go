package service

import (
	"context"
	"fmt"

	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/storage"
)

// MediaService implements the gRPC server for media-related operations.
type MediaService struct {
	storage                                storage.StorageI
	memory.UnimplementedMediaServiceServer // Embed the unimplemented server
}

// NewMediaService creates a new MediaService instance.
func NewMediaService(storage storage.StorageI) *MediaService {
	return &MediaService{
		storage: storage,
	}
}

// GetMediaByID handles the GetMediaByID gRPC request.
func (s *MediaService) GetMediaById(ctx context.Context, req *memory.GetMediaByIdRequest) (*memory.Media, error) {
	media, err := s.storage.Media().GetMediaByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get media by ID: %w", err)
	}

	return media, nil // Return the media message directly
}

// GetAllMedia handles the GetAllMedia gRPC request.
func (s *MediaService) GetAllMedia(ctx context.Context, req *memory.GetAllMediaRequest) (*memory.GetAllMediaResponse, error) {
	mediaList, err := s.storage.Media().GetAllMedia(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all media: %w", err)
	}

	return &memory.GetAllMediaResponse{
		Media: mediaList,
	}, nil
}

// DeleteMedia handles the DeleteMedia gRPC request.
func (s *MediaService) DeleteMedia(ctx context.Context, req *memory.DeleteMediaRequest) (*memory.DeleteMediaResponse, error) {
	if err := s.storage.Media().DeleteMedia(ctx, req.Id); err != nil {
		return nil, fmt.Errorf("failed to delete media: %w", err)
	}

	return &memory.DeleteMediaResponse{}, nil
}
