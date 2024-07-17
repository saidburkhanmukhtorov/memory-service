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

type CommentRepo struct {
	db *pgx.Conn
}

func NewCommentRepo(db *pgx.Conn) *CommentRepo {
	return &CommentRepo{
		db: db,
	}
}

func (r *CommentRepo) CreateComment(ctx context.Context, comment *models.CreateCommentModel) (string, error) {
	if comment.ID == "" {
		comment.ID = uuid.NewString()
	}
	query := `
		INSERT INTO comments (
			id, 
			memory_id,
			user_id,
			content,
			created_at
		) VALUES (
			$1, $2, $3, $4, NOW()
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		comment.ID,
		comment.MemoryID,
		comment.UserID,
		comment.Content,
	).Scan(&comment.ID)

	if err != nil {
		return "", err
	}

	return comment.ID, nil
}

func (r *CommentRepo) GetCommentByID(ctx context.Context, id string) (*memory.Comment, error) {
	var (
		commentModel memory.Comment
		created_at   sql.NullTime
	)
	query := `
		SELECT 
			id,
			memory_id,
			user_id,
			content,
			created_at
		FROM comments
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&commentModel.Id,
		&commentModel.MemoryId,
		&commentModel.UserId,
		&commentModel.Content,
		&created_at,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}
	commentModel.CreatedAt = helper.DateToString(created_at)
	return &commentModel, nil
}

func (r *CommentRepo) GetAllComments(ctx context.Context, req *memory.GetAllCommentsRequest) ([]*memory.Comment, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT 
			id,
			memory_id,
			user_id,
			content,
			created_at
		FROM 
			comments
		WHERE 1=1
	`

	filter := ""

	if req.MemoryId != "" {
		filter += fmt.Sprintf(" AND memory_id = $%d", count)
		args = append(args, req.MemoryId)
		count++
	}

	if req.UserId != "" {
		filter += fmt.Sprintf(" AND user_id = $%d", count)
		args = append(args, req.UserId)
		count++
	}

	if req.Content != "" {
		filter += fmt.Sprintf(" AND content LIKE $%d", count)
		args = append(args, "%"+req.Content+"%")
		count++
	}

	query += filter

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commentList []*memory.Comment

	for rows.Next() {
		var (
			commentModel memory.Comment
			created_at   sql.NullTime
		)
		err = rows.Scan(
			&commentModel.Id,
			&commentModel.MemoryId,
			&commentModel.UserId,
			&commentModel.Content,
			&created_at,
		)
		if err != nil {
			return nil, err
		}
		commentModel.CreatedAt = helper.DateToString(created_at)
		commentList = append(commentList, &commentModel)
	}

	return commentList, nil
}

func (r *CommentRepo) UpdateComment(ctx context.Context, comment *models.UpdateCommentModel) error {
	query := `
		UPDATE comments
		SET 
			memory_id = $1,
			user_id = $2,
			content = $3,
			created_at = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(ctx, query,
		comment.MemoryID,
		comment.UserID,
		comment.Content,
		comment.Created,
		comment.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *CommentRepo) PatchComment(ctx context.Context, comment *models.PatchCommentModel) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE comments
		SET 
	`

	filter := ""

	if comment.MemoryID != nil {
		filter += fmt.Sprintf(" memory_id = $%d, ", count)
		args = append(args, *comment.MemoryID)
		count++
	}

	if comment.UserID != nil {
		filter += fmt.Sprintf(" user_id = $%d, ", count)
		args = append(args, *comment.UserID)
		count++
	}

	if comment.Content != nil {
		filter += fmt.Sprintf(" content = $%d, ", count)
		args = append(args, *comment.Content)
		count++
	}

	if comment.Created != nil {
		filter += fmt.Sprintf(" created_at = $%d, ", count)
		args = append(args, *comment.Created)
		count++
	}

	if filter == "" {
		return fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d", count)
	args = append(args, comment.ID)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *CommentRepo) DeleteComment(ctx context.Context, id string) error {
	query := `
		DELETE FROM comments
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
