package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/models"
	"github.com/time_capsule/memory-service/storage/postgres"
)

func createDBConnection(t *testing.T) *pgx.Conn {

	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"postgres",
	)

	// Connecting to postgres
	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		t.Fatalf("Unable to connect to database: %v", err)
	}
	return db
}

func TestMemoryRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	memoryRepo := postgres.NewMemoryRepo(db)

	t.Run("CreateMemory", func(t *testing.T) {
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory",
			Description: "This is a test memory.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}

		createdID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		// Cleanup
		defer deleteMemory(t, db, createdID)
	})

	t.Run("GetMemoryByID", func(t *testing.T) {
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory",
			Description: "This is a test memory.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}

		createdID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		memory, err := memoryRepo.GetMemoryByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.NotNil(t, memory)
		assert.Equal(t, createdID, memory.Id)

		// Cleanup
		defer deleteMemory(t, db, createdID)
	})

	t.Run("GetAllMemories", func(t *testing.T) {
		// Create some test memories
		createMemoryModel1 := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory 1",
			Description: "This is test memory 1.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		createdID1, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel1)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID1)

		createMemoryModel2 := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory 2",
			Description: "This is test memory 2.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		createdID2, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel2)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID2)

		// Test GetAllMemories
		memories, err := memoryRepo.GetAllMemories(context.Background(), &memory.GetAllMemoriesRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(memories), 2) // At least 2 memories should be returned

		// Cleanup
		defer deleteMemory(t, db, createdID1)
		defer deleteMemory(t, db, createdID2)
	})

	t.Run("UpdateMemory", func(t *testing.T) {
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory",
			Description: "This is a test memory.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}

		createdID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		updateMemoryModel := &models.UpdateMemoryModel{
			ID:          createdID,
			UserID:      createMemoryModel.UserID,
			Title:       "Updated Memory Title",
			Description: "Updated memory description.",
			Date:        time.Now().Add(time.Hour * 24), // Add a day
			Tags:        []string{"updated", "tags"},
			Latitude:    35.0522,
			Longitude:   -119.2437,
			PlaceName:   "San Francisco",
			Privacy:     "private",
		}

		err = memoryRepo.UpdateMemory(context.Background(), updateMemoryModel)
		assert.NoError(t, err)

		updatedMemory, err := memoryRepo.GetMemoryByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, updateMemoryModel.Title, updatedMemory.Title)
		assert.Equal(t, updateMemoryModel.Description, updatedMemory.Description)
		// ... (Assert other updated fields)

		// Cleanup
		defer deleteMemory(t, db, createdID)
	})

	t.Run("PatchMemory", func(t *testing.T) {
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory",
			Description: "This is a test memory.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}

		createdID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		patchMemoryModel := &models.PatchMemoryModel{
			ID:    createdID,
			Title: &createMemoryModel.Title,
		}

		err = memoryRepo.PatchMemory(context.Background(), patchMemoryModel)
		assert.NoError(t, err)

		updatedMemory, err := memoryRepo.GetMemoryByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, createMemoryModel.Title, updatedMemory.Title) // Title should be patched

		// Cleanup
		defer deleteMemory(t, db, createdID)
	})

	t.Run("DeleteMemory", func(t *testing.T) {
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory",
			Description: "This is a test memory.",
			Date:        time.Now(),
			Tags:        []string{"test", "memory"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}

		createdID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		err = memoryRepo.DeleteMemory(context.Background(), createdID)
		assert.NoError(t, err)

		_, err = memoryRepo.GetMemoryByID(context.Background(), createdID)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Memory should not be found
	})
}

// Helper functions for cleanup
func deleteMemory(t *testing.T, db *pgx.Conn, id string) {
	_, err := db.Exec(context.Background(), "DELETE FROM memories WHERE id = $1", id)
	assert.NoError(t, err)
}

func deleteMedia(t *testing.T, db *pgx.Conn, id string) {
	_, err := db.Exec(context.Background(), "DELETE FROM media WHERE id = $1", id)
	assert.NoError(t, err)
}

func deleteComment(t *testing.T, db *pgx.Conn, id string) {
	_, err := db.Exec(context.Background(), "DELETE FROM comments WHERE id = $1", id)
	assert.NoError(t, err)
}
