package test

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/memory-service/config"
	"github.com/time_capsule/memory-service/storage"
	"github.com/time_capsule/memory-service/storage/postgres"
)

// Storage implements the storage.StorageI interface for PostgreSQL.
type Storage struct {
	db       *pgx.Conn
	MemoryS  storage.MemoryI
	MediaS   storage.MediaI
	CommentS storage.CommentI
}

// NewPostgresStorage creates a new PostgreSQL storage instance.
func NewPostgresStorageTest(cfg config.Config) (storage.StorageI, error) {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"postgres",
	)

	db, err := pgx.Connect(context.Background(), dbCon)
	if err != nil {
		slog.Warn("Unable to connect to database:", err)
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		slog.Warn("Unable to ping database:", err)
		return nil, err
	}

	return &Storage{
		db:       db,
		MemoryS:  postgres.NewMemoryRepo(db),
		MediaS:   postgres.NewMediaRepo(db),
		CommentS: postgres.NewCommentRepo(db),
	}, nil
}

// Memory returns the MemoryI implementation for PostgreSQL.
func (s *Storage) Memory() storage.MemoryI {
	return s.MemoryS
}

// Media returns the MediaI implementation for PostgreSQL.
func (s *Storage) Media() storage.MediaI {
	return s.MediaS
}

// Comment returns the CommentI implementation for PostgreSQL.
func (s *Storage) Comment() storage.CommentI {
	return s.CommentS
}
