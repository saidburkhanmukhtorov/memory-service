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

func TestMediaRepo(t *testing.T) {
	db := createDBConnection(t)
	defer db.Close(context.Background())

	memoryRepo := postgres.NewMemoryRepo(db)
	mediaRepo := postgres.NewMediaRepo(db)

	t.Run("CreateMedia", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Media",
			Description: "This is a test memory for media.",
			Date:        time.Now(),
			Tags:        []string{"test", "media"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createMediaModel := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "image",
			URL:      "https://example.com/image.jpg",
			Created:  time.Now(),
		}

		createdID, err := mediaRepo.CreateMedia(context.Background(), createMediaModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		// Cleanup
		defer deleteMedia(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("GetMediaByID", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Media",
			Description: "This is a test memory for media.",
			Date:        time.Now(),
			Tags:        []string{"test", "media"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createMediaModel := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "image",
			URL:      "https://example.com/image.jpg",
			Created:  time.Now(),
		}

		createdID, err := mediaRepo.CreateMedia(context.Background(), createMediaModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		media, err := mediaRepo.GetMediaByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.NotNil(t, media)
		assert.Equal(t, createdID, media.Id)
		assert.Equal(t, memoryID, media.MemoryId)

		// Cleanup
		defer deleteMedia(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("GetAllMedia", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Media",
			Description: "This is a test memory for media.",
			Date:        time.Now(),
			Tags:        []string{"test", "media"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		// Create two test media
		createMediaModel1 := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "image",
			URL:      "https://example.com/image1.jpg",
			Created:  time.Now(),
		}
		createdID1, err := mediaRepo.CreateMedia(context.Background(), createMediaModel1)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID1)

		createMediaModel2 := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "video",
			URL:      "https://example.com/video1.mp4",
			Created:  time.Now(),
		}
		createdID2, err := mediaRepo.CreateMedia(context.Background(), createMediaModel2)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID2)

		// Test GetAllMedia with no filters
		mediaList, err := mediaRepo.GetAllMedia(context.Background(), &memory.GetAllMediaRequest{})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(mediaList), 2) // At least 2 media should be returned

		// Test GetAllMedia with memoryID filter
		mediaList, err = mediaRepo.GetAllMedia(context.Background(), &memory.GetAllMediaRequest{MemoryId: memoryID})
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(mediaList), 2) // At least 2 media should be returned for this memory

		// Cleanup
		defer deleteMedia(t, db, createdID1)
		defer deleteMedia(t, db, createdID2)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("UpdateMedia", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Media",
			Description: "This is a test memory for media.",
			Date:        time.Now(),
			Tags:        []string{"test", "media"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createMediaModel := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "image",
			URL:      "https://example.com/image.jpg",
			Created:  time.Now(),
		}

		createdID, err := mediaRepo.CreateMedia(context.Background(), createMediaModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		updateMediaModel := &models.UpdateMediaModel{
			ID:       createdID,
			MemoryID: memoryID,
			Type:     "video",
			URL:      "https://example.com/video.mp4",
			Created:  time.Now().Add(time.Hour * 24), // Add a day
		}

		err = mediaRepo.UpdateMedia(context.Background(), updateMediaModel)
		assert.NoError(t, err)

		updatedMedia, err := mediaRepo.GetMediaByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, updateMediaModel.Type, updatedMedia.Type)
		assert.Equal(t, updateMediaModel.URL, updatedMedia.Url)
		// ... (Assert other updated fields)

		// Cleanup
		defer deleteMedia(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("PatchMedia", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Media",
			Description: "This is a test memory for media.",
			Date:        time.Now(),
			Tags:        []string{"test", "media"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createMediaModel := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "image",
			URL:      "https://example.com/image.jpg",
			Created:  time.Now(),
		}

		createdID, err := mediaRepo.CreateMedia(context.Background(), createMediaModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		newURL := "https://example.com/new_image.png"
		patchMediaModel := &models.PatchMediaModel{
			ID:  createdID,
			URL: &newURL,
		}

		err = mediaRepo.PatchMedia(context.Background(), patchMediaModel)
		assert.NoError(t, err)

		updatedMedia, err := mediaRepo.GetMediaByID(context.Background(), createdID)
		assert.NoError(t, err)
		assert.Equal(t, newURL, updatedMedia.Url) // URL should be patched

		// Cleanup
		defer deleteMedia(t, db, createdID)
		defer deleteMemory(t, db, memoryID)
	})

	t.Run("DeleteMedia", func(t *testing.T) {
		// Create a test memory first
		createMemoryModel := &models.CreateMemoryModel{
			UserID:      uuid.New().String(),
			Title:       "Test Memory for Media",
			Description: "This is a test memory for media.",
			Date:        time.Now(),
			Tags:        []string{"test", "media"},
			Latitude:    34.0522,
			Longitude:   -118.2437,
			PlaceName:   "Los Angeles",
			Privacy:     "public",
		}
		memoryID, err := memoryRepo.CreateMemory(context.Background(), createMemoryModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, memoryID)

		createMediaModel := &models.CreateMediaModel{
			MemoryID: memoryID,
			Type:     "image",
			URL:      "https://example.com/image.jpg",
			Created:  time.Now(),
		}

		createdID, err := mediaRepo.CreateMedia(context.Background(), createMediaModel)
		assert.NoError(t, err)
		assert.NotEmpty(t, createdID)

		err = mediaRepo.DeleteMedia(context.Background(), createdID)
		assert.NoError(t, err)

		_, err = mediaRepo.GetMediaByID(context.Background(), createdID)
		assert.ErrorIs(t, err, pgx.ErrNoRows) // Media should not be found

		defer deleteMemory(t, db, memoryID)
	})
}
