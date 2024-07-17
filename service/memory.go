package service

import (
	"context"
	"fmt"

	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/storage"
)

// MemoryService implements the gRPC server for memory-related operations.
type MemoryService struct {
	storage                                 storage.StorageI
	memory.UnimplementedMemoryServiceServer // Embed the unimplemented server
}

// NewMemoryService creates a new MemoryService instance.
func NewMemoryService(storage storage.StorageI) *MemoryService {
	return &MemoryService{
		storage: storage,
	}
}

// GetMemoryByID handles the GetMemoryByID gRPC request.
func (s *MemoryService) GetMemoryByID(ctx context.Context, req *memory.GetMemoryByIdRequest) (*memory.Memory, error) {
	memory, err := s.storage.Memory().GetMemoryByID(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get memory by ID: %w", err)
	}

	return memory, nil
}

// GetAllMemories handles the GetAllMemories gRPC request.
func (s *MemoryService) GetAllMemories(ctx context.Context, req *memory.GetAllMemoriesRequest) (*memory.GetAllMemoriesResponse, error) {
	memories, err := s.storage.Memory().GetAllMemories(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get all memories: %w", err)
	}

	return &memory.GetAllMemoriesResponse{
		Memories: memories,
	}, nil
}

// DeleteMemory handles the DeleteMemory gRPC request.
func (s *MemoryService) DeleteMemory(ctx context.Context, req *memory.DeleteMemoryRequest) (*memory.DeleteMemoryResponse, error) {
	if err := s.storage.Memory().DeleteMemory(ctx, req.Id); err != nil {
		return nil, fmt.Errorf("failed to delete memory: %w", err)
	}

	return &memory.DeleteMemoryResponse{}, nil
}
