package main

import (
	"context"
	"log"
	"net"

	"github.com/time_capsule/memory-service/config"
	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/kafka/consumer"
	"github.com/time_capsule/memory-service/service"
	"github.com/time_capsule/memory-service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	// Initialize PostgreSQL storage
	storage, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}

	// Initialize Kafka consumers
	memoryConsumer := consumer.NewMemoryConsumer(cfg.KafkaBrokers, "memory_topic", storage)
	mediaConsumer := consumer.NewMediaConsumer(cfg.KafkaBrokers, "media_topic", storage)
	commentConsumer := consumer.NewCommentConsumer(cfg.KafkaBrokers, "comment_topic", storage)

	// Start consumers in separate goroutines
	go func() {
		if err := memoryConsumer.Consume(context.Background()); err != nil {
			log.Fatalf("memory consumer error: %v", err)
		}
	}()

	go func() {
		if err := mediaConsumer.Consume(context.Background()); err != nil {
			log.Fatalf("media consumer error: %v", err)
		}
	}()

	go func() {
		if err := commentConsumer.Consume(context.Background()); err != nil {
			log.Fatalf("comment consumer error: %v", err)
		}
	}()

	// Initialize gRPC server
	lis, err := net.Listen("tcp", cfg.HTTPPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	memory.RegisterMemoryServiceServer(s, service.NewMemoryService(storage))
	memory.RegisterMediaServiceServer(s, service.NewMediaService(storage))
	memory.RegisterCommentServiceServer(s, service.NewCommentService(storage))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
