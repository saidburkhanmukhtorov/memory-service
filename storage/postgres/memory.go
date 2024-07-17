package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/time_capsule/memory-service/genproto/memory"
	"github.com/time_capsule/memory-service/helper"
	"github.com/time_capsule/memory-service/models"
)

type MemoryRepo struct {
	db *pgx.Conn
}

func NewMemoryRepo(db *pgx.Conn) *MemoryRepo {
	return &MemoryRepo{
		db: db,
	}
}

func (r *MemoryRepo) CreateMemory(ctx context.Context, memory *models.CreateMemoryModel) (string, error) {
	if memory.ID == "" {
		memory.ID = uuid.NewString()
	}
	query := `
		INSERT INTO memories (
			id,
			user_id,
			title,
			description,
			date,
			tags,
			latitude,
			longitude,
			place_name,
			privacy,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW()
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		memory.ID,
		memory.UserID,
		memory.Title,
		memory.Description,
		memory.Date,
		memory.Tags,
		memory.Latitude,
		memory.Longitude,
		memory.PlaceName,
		memory.Privacy,
	).Scan(&memory.ID)

	if err != nil {
		return "", err
	}

	return memory.ID, nil
}

func (r *MemoryRepo) GetMemoryByID(ctx context.Context, id string) (*memory.Memory, error) {
	var (
		memoryModel memory.Memory
		date        sql.NullTime
		created_at  sql.NullTime
		tags        []string
	)
	query := `
		SELECT 
			id,
			user_id,
			title,
			description,
			date,
			tags,
			latitude,
			longitude,
			place_name,
			privacy,
			created_at
		FROM memories
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&memoryModel.Id,
		&memoryModel.UserId,
		&memoryModel.Title,
		&memoryModel.Description,
		&date,
		&tags,
		&memoryModel.Latitude,
		&memoryModel.Longitude,
		&memoryModel.PlaceName,
		&memoryModel.Privacy,
		&created_at,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	memoryModel.Tags = tags
	memoryModel.Date = helper.DateToString(date)
	memoryModel.CreatedAt = helper.DateToString(created_at)
	return &memoryModel, nil
}

func (r *MemoryRepo) GetAllMemories(ctx context.Context, req *memory.GetAllMemoriesRequest) ([]*memory.Memory, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT
			id,
			user_id,
			title,
			description,
			date,
			tags,
			latitude,
			longitude,
			place_name,
			privacy,
			created_at
		FROM 
			memories
		WHERE 1=1 
	`

	filter := ""

	if req.UserId != "" {
		filter += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.Title != "" {
		filter += fmt.Sprintf(" AND title LIKE $%d", count)
		args = append(args, "%"+req.Title+"%")
		count++
	}

	if req.Description != "" {
		filter += fmt.Sprintf(" AND description LIKE $%d", count)
		args = append(args, "%"+req.Description+"%")
		count++
	}

	if len(req.Tags) > 0 {
		filter += fmt.Sprintf(" AND tags && $%d", count) // && operator for array containment
		args = append(args, req.Tags)
		count++
	}

	if req.StartDate != "" {
		startTime, err := time.Parse(time.RFC3339, req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start time format: %w", err)
		}
		filter += fmt.Sprintf(" AND date >= $%d", count)
		args = append(args, startTime)
		count++
	}

	if req.EndDate != "" {
		endTime, err := time.Parse(time.RFC3339, req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end time format: %w", err)
		}
		filter += fmt.Sprintf(" AND date <= $%d", count)
		args = append(args, endTime)
		count++
	}

	if req.Latitude != 0 {
		filter += fmt.Sprintf(" AND latitude = $%d", count)
		args = append(args, req.Latitude)
		count++
	}

	if req.Longitude != 0 {
		filter += fmt.Sprintf(" AND longitude = $%d", count)
		args = append(args, req.Longitude)
		count++
	}

	if req.PlaceName != "" {
		filter += fmt.Sprintf(" AND place_name LIKE $%d", count)
		args = append(args, "%"+req.PlaceName+"%")
		count++
	}

	if req.Privacy != "" {
		filter += fmt.Sprintf(" AND privacy = $%d", count)
		args = append(args, req.Privacy)
		count++
	}

	query += filter

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memories []*memory.Memory

	for rows.Next() {
		var (
			memoryModel memory.Memory
			date        sql.NullTime
			created_at  sql.NullTime
			tags        []string
		)
		err = rows.Scan(
			&memoryModel.Id,
			&memoryModel.UserId,
			&memoryModel.Title,
			&memoryModel.Description,
			&date,
			&tags,
			&memoryModel.Latitude,
			&memoryModel.Longitude,
			&memoryModel.PlaceName,
			&memoryModel.Privacy,
			&created_at,
		)
		if err != nil {
			return nil, err
		}
		memoryModel.Tags = tags
		memoryModel.Date = helper.DateToString(date)
		memoryModel.CreatedAt = helper.DateToString(created_at)
		memories = append(memories, &memoryModel)
	}

	return memories, nil
}

func (r *MemoryRepo) UpdateMemory(ctx context.Context, memory *models.UpdateMemoryModel) error {
	query := `
		UPDATE memories
		SET 
			user_id = $1,
			title = $2,
			description = $3,
			date = $4,
			tags = $5,
			latitude = $6,
			longitude = $7,
			place_name = $8,
			privacy = $9
		WHERE id = $10
	`

	result, err := r.db.Exec(ctx, query,
		memory.UserID,
		memory.Title,
		memory.Description,
		memory.Date,
		memory.Tags,
		memory.Latitude,
		memory.Longitude,
		memory.PlaceName,
		memory.Privacy,
		memory.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *MemoryRepo) PatchMemory(ctx context.Context, memory *models.PatchMemoryModel) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE memories
		SET 
	`

	filter := ""

	if memory.Title != nil {
		filter += fmt.Sprintf(" title = $%d, ", count)
		args = append(args, *memory.Title)
		count++
	}

	if memory.Description != nil {
		filter += fmt.Sprintf(" description = $%d, ", count)
		args = append(args, *memory.Description)
		count++
	}

	if memory.Date != nil {
		filter += fmt.Sprintf(" date = $%d, ", count)
		args = append(args, *memory.Date)
		count++
	}

	if memory.Tags != nil {
		filter += fmt.Sprintf(" tags = $%d, ", count)
		args = append(args, *memory.Tags)
		count++
	}

	if memory.Latitude != nil {
		filter += fmt.Sprintf(" latitude = $%d, ", count)
		args = append(args, *memory.Latitude)
		count++
	}

	if memory.Longitude != nil {
		filter += fmt.Sprintf(" longitude = $%d, ", count)
		args = append(args, *memory.Longitude)
		count++
	}

	if memory.PlaceName != nil {
		filter += fmt.Sprintf(" place_name = $%d, ", count)
		args = append(args, *memory.PlaceName)
		count++
	}

	if memory.Privacy != nil {
		filter += fmt.Sprintf(" privacy = $%d, ", count)
		args = append(args, *memory.Privacy)
		count++
	}

	if filter == "" {
		return fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d", count)
	args = append(args, memory.ID)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *MemoryRepo) DeleteMemory(ctx context.Context, id string) error {
	query := `
		DELETE FROM memories
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
