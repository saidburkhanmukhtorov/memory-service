package service

import (
	"context"
	"fmt"

	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/storage"
)

// ... (MemoryService and MediaService code from previous responses)

// CommentService implements the gRPC server for comment-related operations.
type CommentService struct {
	storage                                  storage.StorageI
	memory.UnimplementedCommentServiceServer // Embed the unimplemented server
}

// NewCommentService creates a new CommentService instance.
func NewCommentService(storage storage.StorageI) *CommentService {
	return &CommentService{
		storage: storage,
	}
}

// GetCommentByID handles the GetCommentByID gRPC request.
func (s *CommentService) GetCommentByID(ctx context.Context, req *memory.GetCommentByIdRequest) (*memory.Comment, error) {
	comment, err := s.storage.Comment().GetCommentByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get comment by ID: %w", err)
	}

	return comment, nil // Return the comment message directly
}

// GetAllComments handles the GetAllComments gRPC request.
func (s *CommentService) GetAllComments(ctx context.Context, req *memory.GetAllCommentsRequest) (*memory.GetAllCommentsResponse, error) {
	commentList, err := s.storage.Comment().GetAllComments(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all comments: %w", err)
	}

	return &memory.GetAllCommentsResponse{
		Comments: commentList,
	}, nil
}

// DeleteComment handles the DeleteComment gRPC request.
func (s *CommentService) DeleteComment(ctx context.Context, req *memory.DeleteCommentRequest) (*memory.DeleteCommentResponse, error) {
	if err := s.storage.Comment().DeleteComment(ctx, req.Id); err != nil {
		return nil, fmt.Errorf("failed to delete comment: %w", err)
	}

	return &memory.DeleteCommentResponse{}, nil
}
