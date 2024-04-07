package services

import "time"

type Comment struct {
	ID          string
	Name        string
	Subject     string
	CommentText string
	CreatedAt   time.Time
}