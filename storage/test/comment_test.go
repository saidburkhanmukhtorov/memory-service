package test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/models"
	"github.com/time_capsule/memory-service/storage/postgres"
)

func TestCommentRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	memoryRepo := postgres.NewMemoryRepo(db)
	commentRepo := postgres.NewCommentRepo(db)

	t.Run("CreateComment", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Comment",
			Description: "This is a test memory for comment.",
			Date:        time.Now(),
			Tags:        []string{"test", "comment"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createCommentModel := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is a test comment.",
			Created:  time.Now(),
		}

		createdID, err := commentRepo.CreateComment(context.Background(), createCommentModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		// Cleanup
		defer deleteComment(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("GetCommentByID", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Comment",
			Description: "This is a test memory for comment.",
			Date:        time.Now(),
			Tags:        []string{"test", "comment"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createCommentModel := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is a test comment.",
			Created:  time.Now(),
		}

		createdID, err := commentRepo.CreateComment(context.Background(), createCommentModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		comment, err := commentRepo.GetCommentByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.NotNil(t, comment)
		assert.Equal(t, createdID, comment.Id)
		assert.Equal(t, memoryID, comment.MemoryId)

		// Cleanup
		defer deleteComment(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("GetAllComments", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Comment",
			Description: "This is a test memory for comment.",
			Date:        time.Now(),
			Tags:        []string{"test", "comment"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		// Create two test comments
		createCommentModel1 := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is test comment 1.",
			Created:  time.Now(),
		}
		createdID1, err := commentRepo.CreateComment(context.Background(), createCommentModel1)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID1)

		createCommentModel2 := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is test comment 2.",
			Created:  time.Now(),
		}
		createdID2, err := commentRepo.CreateComment(context.Background(), createCommentModel2)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID2)

		// Test GetAllComments with no filters
		commentList, err := commentRepo.GetAllComments(context.Background(), &memory.GetAllCommentsRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(commentList), 2) // At least 2 comments should be returned

		// Test GetAllComments with memoryID filter
		commentList, err = commentRepo.GetAllComments(context.Background(), &memory.GetAllCommentsRequest{MemoryId: memoryID})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(commentList), 2) // At least 2 comments should be returned for this memory

		// Cleanup
		defer deleteComment(t, db, createdID1)
		defer deleteComment(t, db, createdID2)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("UpdateComment", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Comment",
			Description: "This is a test memory for comment.",
			Date:        time.Now(),
			Tags:        []string{"test", "comment"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createCommentModel := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is a test comment.",
			Created:  time.Now(),
		}

		createdID, err := commentRepo.CreateComment(context.Background(), createCommentModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		updateCommentModel := &models.UpdateCommentModel{
			ID:       createdID,
			MemoryID: memoryID,
			UserID:   createCommentModel.UserID,
			Content:  "This is an updated comment.",
			Created:  time.Now().Add(time.Hour * 24), // Add a day
		}

		err = commentRepo.UpdateComment(context.Background(), updateCommentModel)
		assert.NoError(t, err)

		updatedComment, err := commentRepo.GetCommentByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, updateCommentModel.Content, updatedComment.Content)
		// ... (Assert other updated fields)

		// Cleanup
		defer deleteComment(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("PatchComment", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Comment",
			Description: "This is a test memory for comment.",
			Date:        time.Now(),
			Tags:        []string{"test", "comment"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createCommentModel := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is a test comment.",
			Created:  time.Now(),
		}

		createdID, err := commentRepo.CreateComment(context.Background(), createCommentModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		newContent := "This is a patched comment."
		patchCommentModel := &models.PatchCommentModel{
			ID:      createdID,
			Content: &newContent,
		}

		err = commentRepo.PatchComment(context.Background(), patchCommentModel)
		assert.NoError(t, err)

		updatedComment, err := commentRepo.GetCommentByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, newContent, updatedComment.Content) // Content should be patched

		// Cleanup
		defer deleteComment(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("DeleteComment", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Comment",
			Description: "This is a test memory for comment.",
			Date:        time.Now(),
			Tags:        []string{"test", "comment"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createCommentModel := &models.CreateCommentModel{
			MemoryID: memoryID,
			UserID:   uuid.New().String(),
			Content:  "This is a test comment.",
			Created:  time.Now(),
		}

		createdID, err := commentRepo.CreateComment(context.Background(), createCommentModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		err = commentRepo.DeleteComment(context.Background(), createdID)
		assert.NoError(t, err)

		_, err = commentRepo.GetCommentByID(context.Background(), createdID)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Comment should not be found

		// Cleanup (memory) - Comment should be already deleted
		defer deleteMemory(t, db, memoryID)
	})
}
