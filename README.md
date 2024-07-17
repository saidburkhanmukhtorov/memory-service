# TimeCapsule Memory Service

This repository contains the Memory service for the TimeCapsule application, a microservice-based system for managing personal memories, media, and comments.

## Architecture

The Memory service is built using the following technologies:

- **gRPC:** For synchronous communication between clients and the service.
- **Kafka:** For asynchronous updates and data persistence.
- **PostgreSQL:** For persistent storage of memories, media, and comments.
- **Docker:** For containerization and deployment.

## Project Structure

.
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│ └── main.go
├── config
│ ├── config.go
│ └── logger
│ └── logger.go
├── docker-compose.yaml
├── genproto
│ └── memory
│ ├── comment.pb.go
│ ├── comment_grpc.pb.go
│ ├── media.pb.go
│ ├── media_grpc.pb.go
│ ├── memory.pb.go
│ └── memory_grpc.pb.go
├── go.mod
├── go.sum
├── helper
│ └── helper.go
├── kafka
│ ├── consermer_test
│ │ └── memory_consumer_test.go
│ └── consumer
│ ├── comment_consumer.go
│ ├── media_consumer.go
│ └── memory_consumer.go
├── models
│ ├── comment.go
│ ├── media.go
│ └── memory.go
├── service
│ ├── comment.go
│ ├── media.go
│ └── memory.go
├── storage
│ ├── postgres
│ │ ├── comment.go
│ │ ├── media.go
│ │ ├── memory.go
│ │ └── postgres.go
│ ├── storage.go
│ └── test
│ ├── comment_test.go
│ ├── media_test.go
│ └── memory_test.go
└── submodule-for-timecapsule
├── README.md
├── memory_models
│ ├── comment.go
│ ├── media.go
│ └── memory.go
├── memory_service
│ ├── comment.proto
│ ├── media.proto
│ └── memory.proto
├── timeline_models
│ ├── custom_event.go
│ ├── historical_event.go
│ └── milestone.go
└── timeline_service
├── custom_event.proto
├── historical_event.proto
└── milestone.proto

## Getting Started

1. **Prerequisites:**

   - Docker
   - Docker Compose
   - Go (version 1.22.5 or later)
   - PostgreSQL
   - Kafka

2. **Set up Environment Variables:**

   - Create a `.env` file in the root of the project with the following environment variables:
     ```
     POSTGRES_HOST=postgres
     POSTGRES_PORT=5432
     POSTGRES_USER=postgres
     POSTGRES_PASSWORD=root
     POSTGRES_DB=memory
     KAFKA_BROKERS=kafka:9092
     ```

3. **Build and Run:**
   - Build the Memory service container:
     ```bash
     docker-compose build memory-service
     ```
   - Run the Memory service container:
     ```bash
     docker-compose up -d
     ```

## Usage

The Memory service provides the following gRPC endpoints:

- **CreateMemory:** Creates a new memory.
- **GetMemoryByID:** Retrieves a memory by its ID.
- **GetAllMemories:** Retrieves all memories.
- **UpdateMemory:** Updates an existing memory.
- **PatchMemory:** Partially updates an existing memory.
- **DeleteMemory:** Deletes a memory by its ID.

## Testing

The project includes a comprehensive test suite for all service methods, storage operations, and Kafka consumers. To run the tests:

```bash
go test ./...
```

**Remember to:**

- **Update the README with specific instructions for your project.**
- **Add any additional information about your project, such as dependencies, configuration, or deployment instructions.**
- **Include a section on how to use the Memory service API.**
- **Consider adding a section on how to contribute to the project.**
- **Choose a suitable license for your project.**

This README file provides a more complete and informative guide for users and contributors to your Memory service project.
