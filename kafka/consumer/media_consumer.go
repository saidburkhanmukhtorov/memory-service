package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/time_capsule/memory-service/models"
	"github.com/time_capsule/memory-service/storage"
)

// MediaConsumer consumes Kafka messages related to media.
type MediaConsumer struct {
	reader  *kafka.Reader
	storage storage.StorageI
}

// NewMediaConsumer creates a new MediaConsumer instance.
func NewMediaConsumer(kafkaBrokers []string, topic string, storage storage.StorageI) *MediaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
		GroupID: "media-group", // Choose a suitable group ID
	})
	return &MediaConsumer{reader: reader, storage: storage}
}

// Consume starts consuming messages from the Kafka topic.
func (c *MediaConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("error fetching message: %w", err)
		}

		var mediaModel models.CreateMediaModel // Or the appropriate model for the message
		if err := json.Unmarshal(msg.Value, &mediaModel); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue // Skip to the next message
		}

		// Process the mediaModel based on the message type (create, update, patch)
		switch string(msg.Key) {
		case "media.create":
			if _, err := c.storage.Media().CreateMedia(ctx, &mediaModel); err != nil {
				log.Printf("error creating media: %v", err)
			}
		case "media.update":
			updateModel := &models.UpdateMediaModel{
				ID:       mediaModel.ID,
				MemoryID: mediaModel.MemoryID,
				Type:     mediaModel.Type,
				URL:      mediaModel.URL,
				Created:  mediaModel.Created,
			}
			if err := c.storage.Media().UpdateMedia(ctx, updateModel); err != nil {
				log.Printf("error updating media: %v", err)
			}
		case "media.patch":
			patchModel := &models.PatchMediaModel{
				ID:       mediaModel.ID,
				MemoryID: &mediaModel.MemoryID,
				Type:     &mediaModel.Type,
				URL:      &mediaModel.URL,
				Created:  &mediaModel.Created,
			}
			if err := c.storage.Media().PatchMedia(ctx, patchModel); err != nil {
				log.Printf("error patching media: %v", err)
			}
		default:
			log.Printf("unknown message key: %s", msg.Key)
		}

		// Commit the message
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("error committing message: %w", err)
		}
	}
}
