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

// MemoryConsumer consumes Kafka messages related to memories.
type MemoryConsumer struct {
	reader  *kafka.Reader
	storage storage.StorageI
}

// NewMemoryConsumer creates a new MemoryConsumer instance.
func NewMemoryConsumer(kafkaBrokers []string, topic string, storage storage.StorageI) *MemoryConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
		GroupID: "memory-group",
	})
	return &MemoryConsumer{reader: reader, storage: storage}
}

// Consume starts consuming messages from the Kafka topic.
func (c *MemoryConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("error fetching message: %w", err)
		}
		var memoryModel models.CreateMemoryModel // Or the appropriate model for the message
		if err := json.Unmarshal(msg.Value, &memoryModel); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue // Skip to the next message
		}

		// Process the memoryModel based on the message type (create, update, patch)
		switch string(msg.Key) {
		case "memory.create":
			if _, err := c.storage.Memory().CreateMemory(ctx, &memoryModel); err != nil {
				log.Printf("error creating memory: %v", err)
			}

		case "memory.update":
			updateModel := &models.UpdateMemoryModel{
				ID:          memoryModel.ID, // Assuming ID is part of the message
				UserID:      memoryModel.UserID,
				Title:       memoryModel.Title,
				Description: memoryModel.Description,
				Date:        memoryModel.Date,
				Tags:        memoryModel.Tags,
				Latitude:    memoryModel.Latitude,
				Longitude:   memoryModel.Longitude,
				PlaceName:   memoryModel.PlaceName,
				Privacy:     memoryModel.Privacy,
			}
			if err := c.storage.Memory().UpdateMemory(ctx, updateModel); err != nil {
				log.Printf("error updating memory: %v", err)
			}
		case "memory.patch":
			patchModel := &models.PatchMemoryModel{
				ID:          memoryModel.ID, // Assuming ID is part of the message
				Title:       &memoryModel.Title,
				Description: &memoryModel.Description,
				Date:        &memoryModel.Date,
				Tags:        &memoryModel.Tags,
				Latitude:    &memoryModel.Latitude,
				Longitude:   &memoryModel.Longitude,
				PlaceName:   &memoryModel.PlaceName,
				Privacy:     &memoryModel.Privacy,
			}
			if err := c.storage.Memory().PatchMemory(ctx, patchModel); err != nil {
				log.Printf("error patching memory: %v", err)
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
