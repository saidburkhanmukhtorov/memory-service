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

// CommentConsumer consumes Kafka messages related to comments.
type CommentConsumer struct {
	reader  *kafka.Reader
	storage storage.StorageI
}

// NewCommentConsumer creates a new CommentConsumer instance.
func NewCommentConsumer(kafkaBrokers []string, topic string, storage storage.StorageI) *CommentConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
		GroupID: "comment-group", // Choose a suitable group ID
	})
	return &CommentConsumer{reader: reader, storage: storage}
}

// Consume starts consuming messages from the Kafka topic.
func (c *CommentConsumer) Consume(ctx context.Context) error {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			return fmt.Errorf("error fetching message: %w", err)
		}

		var commentModel models.CreateCommentModel // Or the appropriate model for the message
		if err := json.Unmarshal(msg.Value, &commentModel); err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue // Skip to the next message
		}

		// Process the commentModel based on the message type (create, update, patch)
		switch string(msg.Key) {
		case "comment.create":
			if _, err := c.storage.Comment().CreateComment(ctx, &commentModel); err != nil {
				log.Printf("error creating comment: %v", err)
			}
		case "comment.update":
			updateModel := &models.UpdateCommentModel{
				ID:       commentModel.ID,
				MemoryID: commentModel.MemoryID,
				UserID:   commentModel.UserID,
				Content:  commentModel.Content,
				Created:  commentModel.Created,
			}
			if err := c.storage.Comment().UpdateComment(ctx, updateModel); err != nil {
				log.Printf("error updating comment: %v", err)
			}
		case "comment.patch":
			patchModel := &models.PatchCommentModel{
				ID:       commentModel.ID,
				MemoryID: &commentModel.MemoryID,
				UserID:   &commentModel.UserID,
				Content:  &commentModel.Content,
				Created:  &commentModel.Created,
			}
			if err := c.storage.Comment().PatchComment(ctx, patchModel); err != nil {
				log.Printf("error patching comment: %v", err)
			}
		default:
			log.Printf("unknown message key: %s", msg.Key)
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("error committing message: %w", err)
		}
	}
}
