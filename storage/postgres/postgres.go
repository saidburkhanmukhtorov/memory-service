package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/memory-service/config"
	"github.com/time_capsule/memory-service/storage"
)

// Storage implements the storage.StorageI interface for PostgreSQL.
type Storage struct {
	db       *pgx.Conn
	MemoryS  storage.MemoryI
	MediaS   storage.MediaI
	CommentS storage.CommentI
}

// NewPostgresStorage creates a new PostgreSQL storage instance.
func NewPostgresStorage(cfg config.Config) (storage.StorageI, error) {
	dbCon := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
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
		MemoryS:  NewMemoryRepo(db),
		MediaS:   NewMediaRepo(db),
		CommentS: NewCommentRepo(db),
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
