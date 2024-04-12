package services

import (
	"context"
	"time"

	"github.com/awoelf/go-web-app/utils"
)

type Comment struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Subject     string    `json:"subject"`
	CommentText string    `json:"commentText"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func setTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), dbTimeout)
}

// GetAllComments
func (c *Comment) GetAllComments() ([]*Comment, error) {
	ctx, cancel := setTimeout()
	defer cancel()

	query := `SELECT id, name, subject, commentText, createdAt, updatedAt FROM comments ORDER BY createdAt DESC`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var comments []*Comment

	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Name,
			&comment.Subject,
			&comment.CommentText,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

// GetComment
func (c *Comment) GetComment(id string) (*Comment, error) {
	ctx, cancel := setTimeout()
	defer cancel()

	query := `
		SELECT id, name, subject, commentText, createdAt, updatedAt 
		FROM comments
		WHERE id = $1
	`

	var comment Comment

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&comment.ID,
		&comment.Name,
		&comment.Subject,
		&comment.CommentText,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// UpdateComment
func (c *Comment) UpdateComment(id string, comment Comment) (*Comment, error) {
	ctx, cancel := setTimeout()
	defer cancel()

	query := `
		UPDATE comments
		SET
			name = $2
			subject = $3
			commentText = $4
			updatedAt = $5
		WHERE id = $1
		returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		id,
		comment.Name,
		comment.Subject,
		comment.CommentText,
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// CreateComment
func (c *Comment) CreateComment(comment Comment) (*Comment, error) {
	ctx, cancel := setTimeout()
	defer cancel()

	query := `
		INSERT INTO comments (id, name, subject, commentText, createdAt, updatedAt)
		VALUES($1, $2, $3, $4, $5, $6) returning *
	`

	comment.ID = utils.GenerateId()

	_, err := db.ExecContext(
		ctx,
		query,
		comment.ID,
		comment.Name,
		comment.Subject,
		comment.CommentText,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment
func (c *Comment) DeleteComment(id string) error {
	ctx, cancel := setTimeout()
	defer cancel()

	query := `DELETE FROM comments WHERE id = $1`

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
