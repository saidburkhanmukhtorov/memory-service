package consumer_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/time_capsule/memory-service/config"
	"github.com/time_capsule/memory-service/kafka/consumer"
	"github.com/time_capsule/memory-service/models"
	"github.com/time_capsule/memory-service/storage/test"
)

func TestMemoryConsumer(t *testing.T) {
	cfg := config.Load()

	// Create a test topic
	topic := "test-memory-topic"
	createTopic(t, []string{"localhost:9092"}, topic)
	defer deleteTopic(t, []string{"localhost:9092"}, topic)

	storage, err := test.NewPostgresStorageTest(cfg)
	if err != nil {
		t.Fatalf("failed to initialize storage: %v", err)
	}

	// Create a test memory model
	memoryModel := &models.CreateMemoryModel{
		ID:          uuid.NewString(),
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

	// Produce a message to the Kafka topic
	produceMessage(t, []string{"localhost:9092"}, topic, "memory.create", memoryModel)

	// Create a MemoryConsumer with the actual storage
	consumer := consumer.NewMemoryConsumer([]string{"localhost:9092"}, topic, storage)

	// Consume the message
	go func() {
		if err := consumer.Consume(context.Background()); err != nil {
			t.Errorf("Error consuming message: %v", err)
		}
	}()

	// Wait for the message to be consumed (adjust timeout as needed)
	time.Sleep(time.Second * 2)

	// Retrieve the created memory from the database
	createdMemory, err := storage.Memory().GetMemoryByID(context.Background(), memoryModel.ID)
	assert.NoError(t, err)
	assert.NotNil(t, createdMemory)

	// Assertions
	assert.Equal(t, memoryModel.UserID, createdMemory.UserId)
	assert.Equal(t, memoryModel.Title, createdMemory.Title)
	assert.Equal(t, memoryModel.Description, createdMemory.Description)
}

// Helper functions to create, delete, and produce messages to a Kafka topic
func createTopic(t *testing.T, brokers []string, topic string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		t.Fatalf("failed to dial leader: %v", err)
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		t.Fatalf("failed to create topic: %v", err)
	}
}

func deleteTopic(t *testing.T, brokers []string, topic string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], topic, 0)
	if err != nil {
		t.Fatalf("failed to dial leader: %v", err)
	}
	defer conn.Close()

	err = conn.DeleteTopics(topic)
	if err != nil {
		t.Fatalf("failed to delete topic: %v", err)
	}
}

func produceMessage(t *testing.T, brokers []string, topic string, key string, message interface{}) {
	w := &kafka.Writer{
		Addr:  kafka.TCP(brokers...),
		Topic: topic,
	}

	value, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("failed to marshal message: %v", err)
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key),
		Value: value,
	})
	if err != nil {
		t.Fatalf("failed to write messages: %v", err)
	}

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close writer: %v", err)
	}
}
