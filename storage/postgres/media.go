package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/helper"
	"github.com/time_capsule/memory-service/models"
)

type MediaRepo struct {
	db *pgx.Conn
}

func NewMediaRepo(db *pgx.Conn) *MediaRepo {
	return &MediaRepo{
		db: db,
	}
}

func (r *MediaRepo) CreateMedia(ctx context.Context, media *models.CreateMediaModel) (string, error) {
	if media.ID == "" {
		media.ID = uuid.NewString()
	}
	query := `
		INSERT INTO media (
			id,
			memory_id,
			type,
			url,
			created_at
		) VALUES (
			$1, $2, $3, $4, NOW()
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		media.ID,
		media.MemoryID,
		media.Type,
		media.URL,
	).Scan(&media.ID)

	if err != nil {
		return "", err
	}

	return media.ID, nil
}

func (r *MediaRepo) GetMediaByID(ctx context.Context, id string) (*memory.Media, error) {
	var (
		mediaModel memory.Media
		created_at sql.NullTime
	)
	query := `
		SELECT 
			id,
			memory_id,
			type,
			url,
			created_at
		FROM media
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&mediaModel.Id,
		&mediaModel.MemoryId,
		&mediaModel.Type,
		&mediaModel.Url,
		&created_at,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}
	mediaModel.CreatedAt = helper.DateToString(created_at)
	return &mediaModel, nil
}

func (r *MediaRepo) GetAllMedia(ctx context.Context, req *memory.GetAllMediaRequest) ([]*memory.Media, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT 
			id,
			memory_id,
			type,
			url,
			created_at
		FROM 
			media
		WHERE 1=1
	`

	filter := ""

	if req.MemoryId != "" {
		filter += fmt.Sprintf(" AND memory_id = $%d", count)
		args = append(args, req.MemoryId)
		count++
	}

	if len(req.Types) > 0 {
		filter += " AND type IN ("
		for i, t := range req.Types {
			if i > 0 {
				filter += ", "
			}
			filter += fmt.Sprintf("$%d", count)
			args = append(args, t)
			count++
		}
		filter += ")"
	}

	query += filter

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mediaList []*memory.Media

	for rows.Next() {
		var (
			mediaModel memory.Media
			created_at sql.NullTime
		)
		err = rows.Scan(
			&mediaModel.Id,
			&mediaModel.MemoryId,
			&mediaModel.Type,
			&mediaModel.Url,
			&created_at,
		)
		if err != nil {
			return nil, err
		}
		mediaModel.CreatedAt = helper.DateToString(created_at)
		mediaList = append(mediaList, &mediaModel)
	}

	return mediaList, nil
}

func (r *MediaRepo) UpdateMedia(ctx context.Context, media *models.UpdateMediaModel) error {
	query := `
		UPDATE media
		SET 
			memory_id = $1,
			type = $2,
			url = $3,
			created_at = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(ctx, query,
		media.MemoryID,
		media.Type,
		media.URL,
		media.Created,
		media.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *MediaRepo) PatchMedia(ctx context.Context, media *models.PatchMediaModel) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE media
		SET 
	`

	filter := ""

	if media.MemoryID != nil {
		filter += fmt.Sprintf(" memory_id = $%d, ", count)
		args = append(args, *media.MemoryID)
		count++
	}

	if media.Type != nil {
		filter += fmt.Sprintf(" type = $%d, ", count)
		args = append(args, *media.Type)
		count++
	}

	if media.URL != nil {
		filter += fmt.Sprintf(" url = $%d, ", count)
		args = append(args, *media.URL)
		count++
	}

	if media.Created != nil {
		filter += fmt.Sprintf(" created_at = $%d, ", count)
		args = append(args, *media.Created)
		count++
	}

	if filter == "" {
		return fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d", count)
	args = append(args, media.ID)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *MediaRepo) DeleteMedia(ctx context.Context, id string) error {
	query := `
		DELETE FROM media
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
